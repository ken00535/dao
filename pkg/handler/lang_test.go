package lodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEqualTrue(t *testing.T) {
	team1 := make(map[string]personType)
	team1["US"] = personType{Name: "Ken", Gender: "Male"}
	team1["JP"] = personType{Name: "Akemi", Gender: "Female"}
	team2 := make(map[string]personType)
	team2["US"] = personType{Name: "Ken", Gender: "Male"}
	team2["JP"] = personType{Name: "Akemi", Gender: "Female"}
	expect := true
	actual := IsEqual(team1, team2)
	assert.Equal(t, expect, actual)
}

func TestIsEqualFalse(t *testing.T) {
	team1 := make(map[string]personType)
	team1["US"] = personType{Name: "Kenny", Gender: "Male"}
	team1["JP"] = personType{Name: "Akemi", Gender: "Female"}
	team2 := make(map[string]personType)
	team2["US"] = personType{Name: "Ken", Gender: "Male"}
	team2["JP"] = personType{Name: "Akemi", Gender: "Female"}
	expect := false
	actual := IsEqual(team1, team2)
	assert.Equal(t, expect, actual)
}
