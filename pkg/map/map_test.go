package maphandler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name   string
	Gender string
}

func TestFindKey(t *testing.T) {
	team := make(map[string]Person)
	team["US"] = Person{Name: "Ken", Gender: "Male"}
	team["JP"] = Person{Name: "Akemi", Gender: "Female"}
	expect := "JP"
	var actual string
	FindKey(team, &actual, func(person interface{}) bool {
		return person.(Person).Name == "Akemi"
	})
	assert.Equal(t, expect, actual)
}
