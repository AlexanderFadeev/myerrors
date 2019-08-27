package myerrors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeferWrapd(t *testing.T) {
	err := func() (err error) {
		defer Wrapd(&err, "2")

		return New("1")
	}()

	if assert.NotNil(t, err) {
		assert.Equal(t, "2: 1", err.Error())
	}
}

func TestDeferWrapdNil(t *testing.T) {
	err := func() (err error) {
		defer Wrapd(&err, "x")

		return nil
	}()

	assert.Nil(t, err)
}

func TestDeferWrapfd(t *testing.T) {
	err := func() (err error) {
		defer Wrapfd(&err, "%d", 2)

		return New("1")
	}()

	if assert.NotNil(t, err) {
		assert.Equal(t, "2: 1", err.Error())
	}
}

func TestDeferWrapfdNil(t *testing.T) {
	err := func() (err error) {
		defer Wrapfd(&err, "%s", "x")

		return nil
	}()

	assert.Nil(t, err)
}

func TestCalld(t *testing.T) {
	err := func() (err error) {
		defer Calld(&err, func() error {
			return New("X")
		})
		return
	}()

	if assert.NotNil(t, err) {
		assert.Equal(t, "X", err.Error())
	}
}

func TestCalldNotOverride(t *testing.T) {
	err := func() (err error) {
		defer Calld(&err, func() error {
			return New("X")
		})
		return New("Y")
	}()

	if assert.NotNil(t, err) {
		assert.Equal(t, "Y", err.Error())
	}
}

func TestCalldNil(t *testing.T) {
	err := func() (err error) {
		defer Calld(&err, func() error {
			return nil
		})
		return New("Y")
	}()

	if assert.NotNil(t, err) {
		assert.Equal(t, "Y", err.Error())
	}
}
