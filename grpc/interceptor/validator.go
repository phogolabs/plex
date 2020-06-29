package interceptor

import (
	"context"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	translator "github.com/go-playground/universal-translator"
	validate "github.com/go-playground/validator/v10"
	validateEn "github.com/go-playground/validator/v10/translations/en"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

// Validatable represents a validatable type
type Validatable interface {
	Validate() error
}

// Validator is the validation interceptor
var Validator = &ValidationHandler{
	Validator:  validate.New(),
	Translator: translator.New(en.New()),
}

func init() {
	translator, _ := Validator.Translator.GetTranslator("en")
	// register the translations
	validateEn.RegisterDefaultTranslations(Validator.Validator, translator)
	// prepare the validator
	Validator.Validator.RegisterTagNameFunc(ExtractValidationFieldName)
}

// ValidationHandler represents a logger
type ValidationHandler struct {
	Validator  *validate.Validate
	Translator *translator.UniversalTranslator
}

// RegisterValidationCtx does the same as RegisterValidation on accepts a FuncCtx validation
// allowing context.Context validation support.
func (h *ValidationHandler) RegisterValidationCtx(tag string, fn validate.FuncCtx, callValidationEvenIfNull ...bool) error {
	return h.Validator.RegisterValidationCtx(tag, fn, callValidationEvenIfNull...)
}

// Unary does unary validation
func (h *ValidationHandler) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
	if validator, ok := req.(Validatable); ok {
		err = validator.Validate()
	} else {

		// validate
		err = h.Validator.StructCtx(ctx, req)
	}

	if err != nil {
		return nil, h.errorf(ctx, err)
	}

	return handler(ctx, req)

}

// Stream does not validate the stream
func (h *ValidationHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, stream)
}

func (h *ValidationHandler) errorf(ctx context.Context, err error) error {
	state, ok := status.FromError(err)
	if ok {
		return state.Err()
	}

	if errs, ok := err.(validate.ValidationErrors); ok {
		state = status.New(codes.InvalidArgument, "unprocessable entity")

		var (
			locale        = h.locale(ctx)
			translator, _ = h.Translator.GetTranslator(locale)
			kv            = make(map[string]interface{}, len(errs))
		)

		// map the error fields
		for _, ferr := range errs {
			// translate the error
			kv[ferr.Field()] = ferr.Translate(translator)
		}

		// prepare the details
		if details, err := structpb.NewStruct(kv); err == nil {
			// add the error as details
			if state, err = state.WithDetails(details); err != nil {
				return err
			}
		}

	} else {
		state = status.New(codes.InvalidArgument, err.Error())
	}

	return state.Err()
}

func (h *ValidationHandler) locale(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if locale := md.Get("Accept-Language"); len(locale) > 0 {
			if lang := locale[0]; lang != "" {
				return lang
			}
		}
	}

	return "en"
}

// ExtractValidationFieldName returns the validated field name
func ExtractValidationFieldName(field reflect.StructField) string {
	tag := field.Tag.Get("json")
	if tag == "" {
		return ""
	}

	name := strings.SplitN(tag, ",", 2)[0]
	if name == "-" {
		return ""
	}

	return name
}
