package http

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/url"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/phogolabs/inflate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	// WithFormMarshaler allows marshaling application/x-www-form-urlencoded requests
	WithFormMarshaler = runtime.WithMarshalerOption(ContentTypeForm,
		&FormMarshaler{},
	)

	// WithJSONMarshaler allows marshaling application/x-www-form-urlencoded requests
	WithJSONMarshaler = runtime.WithMarshalerOption(ContentTypeJSON,
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
		})
)

// FormMarshaler defines a conversion between byte sequence and gRPC payloads / fields.
type FormMarshaler struct{}

// Marshal marshals "v" into byte sequence.
func (m *FormMarshaler) Marshal(v interface{}) ([]byte, error) {
	return nil, status.Error(codes.Unimplemented, "Method Marshal is not supported operation")
}

// Unmarshal unmarshals "data" into "v".
// "v" must be a pointer value.
func (m *FormMarshaler) Unmarshal(data []byte, v interface{}) error {
	decoder := &FormDecoder{Reader: bytes.NewBuffer(data)}
	return decoder.Decode(v)
}

// NewDecoder returns a Decoder which reads byte sequence from "r".
func (m *FormMarshaler) NewDecoder(r io.Reader) runtime.Decoder {
	return &FormDecoder{Reader: r}
}

// NewEncoder returns an Encoder which writes bytes sequence into "w".
func (m *FormMarshaler) NewEncoder(w io.Writer) runtime.Encoder {
	panic(status.Error(codes.Unimplemented, "Method NewEncoder is not supported operation"))
}

// ContentType returns the Content-Type which this marshaler is responsible for.
func (m *FormMarshaler) ContentType(_ interface{}) string {
	return ContentTypeForm
}

// FormDecoder represents a form decoder
type FormDecoder struct {
	Reader io.Reader
}

// Decode decodes the request
func (d *FormDecoder) Decode(v interface{}) error {
	// read the body
	body, err := ioutil.ReadAll(d.Reader)
	if err != nil {
		return err
	}

	// parse the body
	values, err := url.ParseQuery(string(body))
	if err != nil {
		return err
	}

	return inflate.NewFormDecoder(values).Decode(v)
}
