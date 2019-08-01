package myerrors

import (
	"errors"
	pkgerrors "github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() pkgerrors.StackTrace
}

func StackTrace(err error) pkgerrors.StackTrace {
	finder := stackTraceFinder{
		err: err,
	}

	return finder.findStackTrace()
}

type stackTraceFinder struct {
	err          error
	oldestTracer stackTracer
}

func (f *stackTraceFinder) findStackTrace() pkgerrors.StackTrace {
	for f.err != nil {
		f.checkIfTracer(f.err)
		f.checkIfWrapper()

		f.err = errors.Unwrap(f.err)
	}

	if f.oldestTracer == nil {
		return nil
	}

	return f.oldestTracer.StackTrace()
}

func (f *stackTraceFinder) checkIfTracer(err error) {
	if tracer, ok := err.(stackTracer); ok {
		f.oldestTracer = tracer
	}
}

func (f *stackTraceFinder) checkIfWrapper() {
	if wrapper, ok := f.err.(*wrapper); ok {
		f.checkIfTracer(wrapper.error)
	}
}
