package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	func1 := func(i int) int {
		return (i + 1)
	}
	var func2 func() int
	Wrap(1, func1, &func2)
	actual := func2()
	expect := 2
	assert.Equal(t, expect, actual)
}
