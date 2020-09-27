package dao

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name   string
	Gender string
}

func TestFilter(t *testing.T) {
	people := []Person{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	actual := []Person{}
	expect := []Person{{Name: "Ken", Gender: "Male"}}
	Filter(people, &actual, func(person interface{}) bool {
		return person.(Person).Name == "Ken"
	})
	assert.Equal(t, expect, actual)
}

func TestOmit(t *testing.T) {
	people := []Person{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	actual := []Person{}
	expect := []Person{{Name: "Cythia", Gender: "Female"}}
	Omit(people, &actual, func(person interface{}) bool {
		return person.(Person).Name == "Ken"
	})
	assert.Equal(t, expect, actual)
}

func TestForEach(t *testing.T) {
	people := []Person{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	var actual []string
	expect := []string{"Ken", "Cythia"}
	ForEach(people, func(person interface{}) {
		actual = append(actual, person.(Person).Name)
	})
	assert.Equal(t, expect, actual)
}

func TestReduce(t *testing.T) {
	people := []Person{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	expect := Person{Name: "Ken Cythia"}
	actual := Person{}
	Reduce(people, &actual, func(m interface{}, n interface{}) interface{} {
		person := Person{
			Name: m.(Person).Name + " " + n.(Person).Name,
		}
		return person
	})
	assert.Equal(t, expect, actual)
}

func TestMap(t *testing.T) {
	people := []Person{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	expect := []Person{{Name: "ken", Gender: "Male"}, {Name: "cythia", Gender: "Female"}}
	actual := []Person{}
	Map(people, &actual, func(person interface{}) interface{} {
		val := person.(Person)
		val.Name = strings.ToLower(val.Name)
		return val
	})
	assert.Equal(t, expect, actual)
}

func TestFind(t *testing.T) {
	people := []Person{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	expect := Person{Name: "Cythia", Gender: "Female"}
	actual := Person{}
	Find(people, &actual, func(person interface{}) bool {
		return person.(Person).Name == "Cythia"
	})
	assert.Equal(t, expect, actual)
}
