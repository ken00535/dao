package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBind(t *testing.T) {
	multiple := func(num1 int, num2 float32) int {
		product := float32(num1) * num2
		return int(product)
	}
	var bindFunc func(float32) int
	Bind(multiple, &bindFunc, 2)
	actual := bindFunc(3.5)
	expect := 7
	assert.Equal(t, expect, actual)
}
