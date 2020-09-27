package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKey(t *testing.T) {
	team := make(map[string]personType)
	team["US"] = personType{Name: "Ken", Gender: "Male"}
	team["JP"] = personType{Name: "Akemi", Gender: "Female"}
	expect := "JP"
	var actual string
	FindKey(team, &actual, func(person interface{}) bool {
		return person.(personType).Name == "Akemi"
	})
	assert.Equal(t, expect, actual)
}

func TestKeys(t *testing.T) {
	team := make(map[string]personType)
	team["US"] = personType{Name: "Ken", Gender: "Male"}
	team["JP"] = personType{Name: "Akemi", Gender: "Female"}
	expect := []string{"JP", "US"}
	var actual []string
	Keys(team, &actual)
	assert.ElementsMatch(t, expect, actual)
}
