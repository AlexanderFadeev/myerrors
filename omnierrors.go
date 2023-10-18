package omnierrors

import (
	stderrors "errors"

	pkgerrors "github.com/pkg/errors"
)

type (
	StackTrace = pkgerrors.StackTrace
)

var (
	New       = pkgerrors.New
	Errorf    = pkgerrors.Errorf
	Wrap      = pkgerrors.Wrap
	Wrapf     = pkgerrors.Wrapf
	WithStack = pkgerrors.WithStack

	Is     = stderrors.Is
	As     = stderrors.As
	Join   = stderrors.Join
	Unwrap = stderrors.Unwrap
)

// Special error type with no stack trace info
type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

// Unlike Join does not combine errors.
// Returns first non-nil error if present.
// Returns nil othwerwise.
func First(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

func GetStackTrace(err error) StackTrace {
	type tracer interface {
		StackTrace() StackTrace
	}

	var deepest tracer

	for err != nil {
		if tracer, ok := err.(tracer); ok {
			deepest = tracer
		}
		err = Unwrap(err)
	}

	if deepest == nil {
		return nil
	}
	return deepest.StackTrace()
}
