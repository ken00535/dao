package lodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatial(t *testing.T) {
	multiple := func(num1 int, num2 float32) int {
		product := float32(num1) * num2
		return int(product)
	}
	var partialFunc func(float32) int
	Partial(multiple, &partialFunc, 2)
	actual := partialFunc(3.5)
	expect := 7
	assert.Equal(t, expect, actual)
}
