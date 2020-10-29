package lodash

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type personType struct {
	Name   string
	Gender string
}

func TestFilter(t *testing.T) {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	actual := []personType{}
	expect := []personType{{Name: "Ken", Gender: "Male"}}
	Filter(people, &actual, func(person interface{}) bool {
		return person.(personType).Name == "Ken"
	})
	assert.Equal(t, expect, actual)
}

func TestFilterNew(t *testing.T) {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	actual := []personType{}
	expect := []personType{{Name: "Ken", Gender: "Male"}}
	Start(people).Filter(func(person interface{}) bool {
		return person.(personType).Name == "Ken"
	}).End(&actual)
	assert.Equal(t, expect, actual)
}

func TestOmit(t *testing.T) {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	actual := []personType{}
	expect := []personType{{Name: "Cythia", Gender: "Female"}}
	Omit(people, &actual, func(person interface{}) bool {
		return person.(personType).Name == "Ken"
	})
	assert.Equal(t, expect, actual)
}

func TestForEach(t *testing.T) {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	var actual []string
	expect := []string{"Ken", "Cythia"}
	ForEach(people, func(person interface{}) {
		actual = append(actual, person.(personType).Name)
	})
	assert.Equal(t, expect, actual)
}

func TestReduce(t *testing.T) {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	expect := personType{Name: "Ken Cythia"}
	actual := personType{}
	Reduce(people, &actual, func(m interface{}, n interface{}) interface{} {
		person := personType{
			Name: m.(personType).Name + " " + n.(personType).Name,
		}
		return person
	})
	assert.Equal(t, expect, actual)
}

func TestMap(t *testing.T) {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	expect := []personType{{Name: "ken", Gender: "Male"}, {Name: "cythia", Gender: "Female"}}
	actual := []personType{}
	Map(people, &actual, func(person interface{}) interface{} {
		val := person.(personType)
		val.Name = strings.ToLower(val.Name)
		return val
	})
	assert.Equal(t, expect, actual)
}

func TestFind(t *testing.T) {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	expect := personType{Name: "Cythia", Gender: "Female"}
	actual := personType{}
	Find(people, &actual, func(person interface{}) bool {
		return person.(personType).Name == "Cythia"
	})
	assert.Equal(t, expect, actual)
}

func TestUniq(t *testing.T) {
	data := []int{1, 1, 2, 2}
	expect := []int{1, 2}
	var actual []int
	Uniq(data, &actual)
	assert.ElementsMatch(t, expect, actual)
}

func TestDifference(t *testing.T) {
	data := []int{1, 1, 2, 2, 3}
	values := []int{1, 2}
	expect := []int{3}
	var actual []int
	Difference(data, values, &actual)
	assert.ElementsMatch(t, expect, actual)
}
