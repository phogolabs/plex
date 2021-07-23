package proto

import (
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// String stores v in a new string value and returns a pointer to it.
func String(value string) *string {
	if len(value) > 0 {
		return &value
	}

	return nil
}

// GetString returns the string from pointer.
func GetString(value *string) string {
	if value != nil {
		return *value
	}

	return ""
}

// Time stores v in a new string value and returns a pointer to it.
func Time(value time.Time) *time.Time {
	if !value.IsZero() {
		return &value
	}

	return nil
}

// GetTime returns the time from pointer.
func GetTime(value *time.Time) time.Time {
	if value != nil {
		return *value
	}

	return time.Time{}
}

// Timestamp converts a time to timestamp.
func Timestamp(value *time.Time) *timestamppb.Timestamp {
	if value != nil {
		return timestamppb.New(*value)
	}

	return nil
}

// FieldMask creates a field mask that usuaally used to update the field.
func FieldMask(columns ...string) *fieldmaskpb.FieldMask {
	return &fieldmaskpb.FieldMask{
		Paths: columns,
	}
}

// OrderByMask creates a field mask that usuaally used to order by the field.
func OrderByMask(clause ...string) *fieldmaskpb.FieldMask {
	columns := []string{}

	var (
		prefix = []string{"+", "-"}
		suffix = []string{"asc", "desc", "ASC", "DESC"}
	)

	for _, path := range clause {
		for _, part := range strings.Split(path, ",") {
			// trim prefix
			for _, symbol := range prefix {
				part = strings.TrimPrefix(part, symbol)
				part = strings.TrimSpace(part)
			}

			// trim suffix
			for _, symbol := range suffix {
				part = strings.TrimSuffix(part, symbol)
				part = strings.TrimSpace(part)
			}

			if len(part) > 0 {
				// column name
				columns = append(columns, part)
			}
		}
	}

	return &fieldmaskpb.FieldMask{
		Paths: columns,
	}
}

// UnionMask returns the union of all the paths in the input field masks.
func UnionMask(x *fieldmaskpb.FieldMask, columns ...string) *fieldmaskpb.FieldMask {
	y := &fieldmaskpb.FieldMask{
		Paths: columns,
	}

	return fieldmaskpb.Union(x, y)
}

// URLRef returns the type name.
func URLRef(x proto.Message) string {
	var (
		message = x.ProtoReflect().Descriptor()
		schema  = x.ProtoReflect().Descriptor().ParentFile()
	)

	return fmt.Sprintf("https://%v/%v", schema.Package(), message.FullName())
}