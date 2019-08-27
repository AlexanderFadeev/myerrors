package myerrors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeANil(t *testing.T) {
	errA := New("A")
	var errB error

	err := Merge(errA, errB)

	if assert.NotNil(t, err) {
		assert.Equal(t, "A", err.Error())
	}
}

func TestMergeNilB(t *testing.T) {
	var errA error
	errB := New("B")

	err := Merge(errA, errB)

	if assert.NotNil(t, err) {
		assert.Equal(t, "B", err.Error())
	}
}

func TestMergeAB(t *testing.T) {
	errA := New("A")
	errB := New("B")

	err := Merge(errA, errB)

	if assert.NotNil(t, err) {
		assert.Equal(t, "A", err.Error())
	}
}

func TestMergeNilNil(t *testing.T) {
	var errA error
	var errB error

	err := Merge(errA, errB)

	assert.Nil(t, err)
}
