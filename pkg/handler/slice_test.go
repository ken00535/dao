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
	type testCase struct {
		dataReq []personType
		want    []personType
	}
	cases := []testCase{
		{
			dataReq: []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}},
			want:    []personType{{Name: "Ken", Gender: "Male"}},
		},
		{
			dataReq: []personType{{Name: "Alisa", Gender: "Female"}, {Name: "Cythia", Gender: "Female"}},
			want:    []personType{},
		},
	}
	for _, tc := range cases {
		actual := []personType{}
		Start(tc.dataReq).Filter(func(person interface{}) bool {
			return person.(personType).Name == "Ken"
		}).End(&actual)
		assert.Equal(t, tc.want, actual)
	}
}

func TestFind(t *testing.T) {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	expect := personType{Name: "Cythia", Gender: "Female"}
	actual := personType{}
	Start(people).Find(func(person interface{}) bool {
		return person.(personType).Name == "Cythia"
	}).End(&actual)
	assert.Equal(t, expect, actual)
}

func TestOmit(t *testing.T) {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	actual := []personType{}
	expect := []personType{{Name: "Cythia", Gender: "Female"}}
	Start(people).Omit(func(person interface{}) bool {
		return person.(personType).Name == "Ken"
	}).End(&actual)
	assert.Equal(t, expect, actual)
}

func TestMap(t *testing.T) {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	expect := []personType{{Name: "ken", Gender: "Male"}, {Name: "cythia", Gender: "Female"}}
	actual := []personType{}
	Start(people).Map(func(person interface{}) interface{} {
		val := person.(personType)
		val.Name = strings.ToLower(val.Name)
		return val
	}).End(&actual)
	assert.Equal(t, expect, actual)
}

func TestReduce(t *testing.T) {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	expect := personType{Name: "Ken Cythia"}
	actual := personType{}
	Start(people).Reduce(func(m interface{}, n interface{}) interface{} {
		person := personType{
			Name: m.(personType).Name + " " + n.(personType).Name,
		}
		return person
	}).End(&actual)
	assert.Equal(t, expect, actual)
}

func TestUniq(t *testing.T) {
	data := []int{1, 1, 2, 2}
	expect := []int{1, 2}
	var actual []int
	Start(data).Uniq().End(&actual)
	assert.ElementsMatch(t, expect, actual)
}

func TestDifference(t *testing.T) {
	data := []int{1, 1, 2, 2, 3}
	values := []int{1, 2}
	expect := []int{3}
	var actual []int
	Start(data).Difference(values).End(&actual)
	assert.ElementsMatch(t, expect, actual)
}

func TestForEach(t *testing.T) {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	var actual []string
	expect := []string{"Ken", "Cythia"}
	Start(people).ForEach(func(person interface{}) bool {
		actual = append(actual, person.(personType).Name)
		return true
	})
	assert.Equal(t, expect, actual)
}

func TestChain(t *testing.T) {
	type testCase struct {
		dataReq []personType
		want    []personType
	}
	cases := []testCase{
		{
			dataReq: []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}},
			want:    []personType{{Name: "ken", Gender: "Male"}},
		},
		{
			dataReq: []personType{{Name: "Alisa", Gender: "Female"}, {Name: "Cythia", Gender: "Female"}},
			want:    []personType{},
		},
	}
	for _, tc := range cases {
		actual := []personType{}
		Start(tc.dataReq).Filter(func(person interface{}) bool {
			return person.(personType).Name == "Ken"
		}).Map(func(person interface{}) interface{} {
			val := person.(personType)
			val.Name = strings.ToLower(val.Name)
			return val
		}).End(&actual)
		assert.Equal(t, tc.want, actual)
	}
}
