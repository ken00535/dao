package lodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKeyNew(t *testing.T) {
	team := make(map[string]personType)
	team["US"] = personType{Name: "Ken", Gender: "Male"}
	team["JP"] = personType{Name: "Akemi", Gender: "Female"}
	expect := "JP"
	var actual string
	Start(team).FindKey(func(person interface{}) bool {
		return person.(personType).Name == "Akemi"
	}).End(&actual)
	assert.Equal(t, expect, actual)
}

func TestKeysNew(t *testing.T) {
	team := make(map[string]personType)
	team["US"] = personType{Name: "Ken", Gender: "Male"}
	team["JP"] = personType{Name: "Akemi", Gender: "Female"}
	expect := []string{"JP", "US"}
	var actual []string
	Start(team).Keys().End(&actual)
	assert.ElementsMatch(t, expect, actual)
}

func TestValuesNew(t *testing.T) {
	team := make(map[string]personType)
	team["US"] = personType{Name: "Ken", Gender: "Male"}
	team["JP"] = personType{Name: "Akemi", Gender: "Female"}
	expect := []personType{
		{Name: "Ken", Gender: "Male"},
		{Name: "Akemi", Gender: "Female"},
	}
	var actual []personType
	Start(team).Values().End(&actual)
	assert.Equal(t, expect, actual)
}

func TestPickNew(t *testing.T) {
	team := make(map[string]personType)
	team["US"] = personType{Name: "Ken", Gender: "Male"}
	team["JP"] = personType{Name: "Akemi", Gender: "Female"}
	expect := make(map[string]personType)
	expect["US"] = personType{Name: "Ken", Gender: "Male"}
	actual := make(map[string]personType)
	Start(team).Pick(func(person interface{}) bool {
		return person.(personType).Name == "Ken"
	}).End(&actual)
	assert.Equal(t, expect, actual)
}

func TestMergeNew(t *testing.T) {
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
	Start(team1).Merge(team2).End(&actual)
	assert.Equal(t, expect, actual)
}
