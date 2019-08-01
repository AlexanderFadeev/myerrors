package myerrors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackTrace(t *testing.T) {
	err := New("1")
	err = Wrap(err, "2")

	trace := StackTrace(err)
	frame := trace[1]
	fileLine := fmt.Sprintf("%s:%d", frame, frame)

	assert.Equal(t, "stack_test.go:11", fileLine)
}

func TestStackTraceNil(t *testing.T) {
	err := errors.New("1")

	trace := StackTrace(err)

	assert.Nil(t, trace)
}

func TestStackTraceWrapNil(t *testing.T) {
	err := errors.New("1")
	err = Wrap(err, "2")

	trace := StackTrace(err)
	frame := trace[1]
	fileLine := fmt.Sprintf("%s:%d", frame, frame)

	assert.Equal(t, "stack_test.go:31", fileLine)
}
