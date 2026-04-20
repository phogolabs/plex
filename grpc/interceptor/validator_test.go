package interceptor_test

import (
	"testing"

	"github.com/phogolabs/plex/grpc/interceptor"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type member struct {
	Email *wrapperspb.StringValue `json:"email" validate:"omitempty,email"`
	Phone *wrapperspb.StringValue `json:"phone" validate:"omitempty,e164"`
}

func TestValidator_StringValueWrapper(t *testing.T) {
	tests := []struct {
		name    string
		input   *member
		wantErr bool
	}{
		{
			name:    "nil wrappers pass omitempty",
			input:   &member{},
			wantErr: false,
		},
		{
			name: "empty-string wrappers pass omitempty",
			input: &member{
				Email: wrapperspb.String(""),
				Phone: wrapperspb.String(""),
			},
			wantErr: false,
		},
		{
			name: "valid email and e164 phone pass",
			input: &member{
				Email: wrapperspb.String("test@example.com"),
				Phone: wrapperspb.String("+12125551234"),
			},
			wantErr: false,
		},
		{
			name: "invalid email fails",
			input: &member{
				Email: wrapperspb.String("not-an-email"),
			},
			wantErr: true,
		},
		{
			name: "non-e164 phone fails",
			input: &member{
				Phone: wrapperspb.String("212-555-1234"),
			},
			wantErr: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := interceptor.Validator.Validator.Struct(tc.input)
			if tc.wantErr && err == nil {
				t.Fatalf("expected validation error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected validation error: %v", err)
			}
		})
	}
}
