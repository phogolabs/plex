package interceptor

import (
	"context"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	translator "github.com/go-playground/universal-translator"
	validate "github.com/go-playground/validator/v10"
	validateEn "github.com/go-playground/validator/v10/translations/en"
	"github.com/phogolabs/log"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	// prepare the validator
	Validator.Validator.RegisterTagNameFunc(ExtractValidationFieldName)
	// register the translations
	if err := validateEn.RegisterDefaultTranslations(Validator.Validator, translator); err != nil {
		panic(err)
	}
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
	logger := log.GetContext(ctx)

	if validator, ok := req.(Validatable); ok {
		err = validator.Validate()
	} else {

		// validate
		err = h.Validator.StructCtx(ctx, req)
	}

	if err != nil {
		logger.WithError(err).Warn("validation failure")
		return nil, h.errorf(ctx, err)
	}

	return handler(ctx, req)

}

// Stream does not validate the stream
func (h *ValidationHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, stream)
}

func (h *ValidationHandler) errorf(ctx context.Context, err error) error {
	werr, ok := status.FromError(err)
	if ok {
		return werr.Err()
	}

	verr, ok := err.(validate.ValidationErrors)
	if !ok {
		werr = status.New(codes.InvalidArgument, err.Error())
		return werr.Err()
	}

	var (
		locale        = h.locale(ctx)
		translator, _ = h.Translator.GetTranslator(locale)
		details       = &errdetails.BadRequest{}
	)

	// map the error fields
	for _, ferr := range verr {
		violation := &errdetails.BadRequest_FieldViolation{
			Field:       ferr.Field(),
			Description: ferr.Translate(translator),
		}

		details.FieldViolations = append(details.FieldViolations, violation)
	}

	werr, err = status.
		New(codes.InvalidArgument, "Unprocessable entity").
		WithDetails(details)

	if err != nil {
		return err
	}

	return werr.Err()
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
