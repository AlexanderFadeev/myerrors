package myerrors

import pkgerrors "github.com/pkg/errors"

func New(message string) error {
	return pkgerrors.New(message)
}

func Errorf(format string, args ...interface{}) error {
	return pkgerrors.Errorf(format, args...)
}

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}

	return &wrapper{pkgerrors.Wrap(err, msg)}
}

func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return &wrapper{pkgerrors.Wrapf(err, format, args...)}
}

func WithStack(err error) error {
	if err == nil {
		return nil
	}

	return &wrapper{pkgerrors.WithStack(err)}
}
