package lodash

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

func TestValues(t *testing.T) {
	team := make(map[string]personType)
	team["US"] = personType{Name: "Ken", Gender: "Male"}
	team["JP"] = personType{Name: "Akemi", Gender: "Female"}
	expect := []personType{
		personType{Name: "Ken", Gender: "Male"},
		personType{Name: "Akemi", Gender: "Female"},
	}
	var actual []personType
	Values(team, &actual)
	assert.Equal(t, expect, actual)
}

func TestPick(t *testing.T) {
	team := make(map[string]personType)
	team["US"] = personType{Name: "Ken", Gender: "Male"}
	team["JP"] = personType{Name: "Akemi", Gender: "Female"}
	expect := make(map[string]personType)
	expect["US"] = personType{Name: "Ken", Gender: "Male"}
	actual := make(map[string]personType)
	Pick(team, &actual, func(person interface{}) bool {
		return person.(personType).Name == "Ken"
	})
	assert.Equal(t, expect, actual)
}

func TestMerge(t *testing.T) {
	team1 := make(map[string]personType)
	team1["US"] = personType{Name: "Ken", Gender: "Male"}
	team1["JP"] = personType{Name: "Akemi", Gender: "Female"}

	team2 := make(map[string]personType)
	team2["GBR"] = personType{Name: "Andy", Gender: "Male"}

	expect := make(map[string]personType)
	expect["US"] = personType{Name: "Ken", Gender: "Male"}
	expect["JP"] = personType{Name: "Akemi", Gender: "Female"}
	expect["GBR"] = personType{Name: "Andy", Gender: "Male"}

	actual := make(map[string]personType)
	Merge(team1, team2, &actual)
	assert.Equal(t, expect, actual)
}
