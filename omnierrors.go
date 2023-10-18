package omnierrors

import (
	"errors"
	"fmt"
)

func New(msg string) error {
	return errors.New(msg)
}

func Errorf(format string, args ...any) error {
	return fmt.Errorf(format, args...)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target error) bool {
	return errors.As(err, target)
}

func Wrap(err error, msg string) error {
	return Errorf("%s: %w", msg, err)
}

func Wrapf(err error, format string, args ...any) error {
	return Wrap(err, fmt.Sprintf(format, args...))
}

// Unlike errors.Join does not combine errors.
// Returns first non-nil error if present.
// Returns nil othwerwise.
func Join(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
