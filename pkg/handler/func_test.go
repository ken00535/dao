package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	multiple := func(num1 int, num2 float32) int {
		product := float32(num1) * num2
		return int(product)
	}
	var wrapFunc func(float32) int
	Wrap(multiple, &wrapFunc, 2)
	actual := wrapFunc(3.5)
	expect := 7
	assert.Equal(t, expect, actual)
}
