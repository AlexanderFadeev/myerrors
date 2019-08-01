package myerrors

import pkgerrors "github.com/pkg/errors"

func New(message string) error {
	return pkgerrors.New(message)
}

func Errorf(format string, args ...interface{}) error {
	return pkgerrors.Errorf(format, args...)
}

func Wrap(err error, msg string) error {
	return &wrapper{pkgerrors.Wrap(err, msg)}
}

func Wrapf(err error, format string, args ...interface{}) error {
	return &wrapper{pkgerrors.Wrapf(err, format, args...)}
}
