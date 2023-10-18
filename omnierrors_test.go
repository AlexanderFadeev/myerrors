package omnierrors_test

import (
	"fmt"
	"testing"

	"github.com/afadeevz/omnierrors"
	"github.com/stretchr/testify/assert"
)

var errTest1 = omnierrors.Sentinel("1")
var errTest2 = omnierrors.Sentinel("2")

func TestStackTrace(t *testing.T) {
	err := omnierrors.New("1")
	err = omnierrors.Wrap(err, "2")

	trace := omnierrors.GetStackTrace(err)
	frame := trace[0]
	fileLine := fmt.Sprintf("%s:%d", frame, frame)

	assert.Equal(t, "omnierrors_test.go:15", fileLine)
}

func TestStackTraceNil(t *testing.T) {
	trace := omnierrors.GetStackTrace(errTest1)

	assert.Nil(t, trace)
}

func TestStackTraceWrapNil(t *testing.T) {
	err := omnierrors.Wrap(errTest1, "2")

	trace := omnierrors.GetStackTrace(err)
	frame := trace[0]
	fileLine := fmt.Sprintf("%s:%d", frame, frame)

	assert.Equal(t, "omnierrors_test.go:32", fileLine)
}

func TestWithStack(t *testing.T) {
	err := omnierrors.WithStack(errTest1)

	trace := omnierrors.GetStackTrace(err)
	frame := trace[0]
	fileLine := fmt.Sprintf("%s:%d", frame, frame)

	assert.Equal(t, "omnierrors_test.go:42", fileLine)
}

func TestFirstEmpty(t *testing.T) {
	err := omnierrors.First()
	assert.Nil(t, err)
}

func TestFirstNil(t *testing.T) {
	err := omnierrors.First(nil)
	assert.Nil(t, err)
}

func TestFirstOne(t *testing.T) {
	err := omnierrors.First(errTest1)
	assert.Equal(t, errTest1, err)
}

func TestFirstNilOne(t *testing.T) {
	err := omnierrors.First(nil, errTest1)
	assert.Equal(t, errTest1, err)
}

func TestFirstOneNil(t *testing.T) {
	err := omnierrors.First(errTest1, nil)
	assert.Equal(t, errTest1, err)
}

func TestFirstOneTwo(t *testing.T) {
	err := omnierrors.First(errTest1, errTest2)
	assert.Equal(t, errTest1, err)
}
