package omnierrors_test

import (
	"testing"

	"github.com/afadeevz/omnierrors"
	"github.com/stretchr/testify/assert"
)

var (
	errTest1 = omnierrors.New("1")
	errTest2 = omnierrors.New("2")
)

func TestJoinEmpty(t *testing.T) {
	t.Parallel()
	err := omnierrors.Join()
	assert.Nil(t, err)
}

func TestJoinNil(t *testing.T) {
	t.Parallel()
	err := omnierrors.Join(nil)
	assert.Nil(t, err)
}

func TestJoinOne(t *testing.T) {
	t.Parallel()
	err := omnierrors.Join(errTest1)
	assert.Equal(t, errTest1, err)
}

func TestJoinNilOne(t *testing.T) {
	t.Parallel()
	err := omnierrors.Join(nil, errTest1)
	assert.Equal(t, errTest1, err)
}

func TestJoinOneNil(t *testing.T) {
	t.Parallel()
	err := omnierrors.Join(errTest1, nil)
	assert.Equal(t, errTest1, err)
}

func TestJoinOneTwo(t *testing.T) {
	t.Parallel()
	err := omnierrors.Join(errTest1, errTest2)
	assert.Equal(t, errTest1, err)
}
