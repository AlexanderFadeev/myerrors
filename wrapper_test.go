package myerrors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ Wrapper = &wrapper{}

func TestWrap(t *testing.T) {
	err := errors.New("1")
	err = Wrap(err, "2")
	err = Wrapf(err, "%d", 3)

	assert.NotNil(t, err)
	if assert.NotNil(t, err) {
		assert.Equal(t, "3: 2: 1", err.Error())
	}

	u := errors.Unwrap(err)
	u = errors.Unwrap(u)

	if assert.NotNil(t, u) {
		assert.Equal(t, "2: 1", u.Error())
	}

	u = errors.Unwrap(u)
	u = errors.Unwrap(u)

	if assert.NotNil(t, u) {
		assert.Equal(t, "1", u.Error())
	}
}
