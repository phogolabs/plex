package health

import (
	"context"
)

// Checker checks the health
type Checker interface {
	Name() string
	Check(context.Context) error
}

var _ Checker = CheckerFunc(nil)

// CheckerFunc represents a cheker function
type CheckerFunc func(context.Context) error

// Name returns the name of the cheker
func (fn CheckerFunc) Name() string {
	return "func"
}

// Check checks the resource
func (fn CheckerFunc) Check(ctx context.Context) error {
	return fn(ctx)
}
