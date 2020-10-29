package main

import (
	"fmt"
	"strings"

	lo "github.com/ken00535/lodash/pkg/handler"
)

type personType struct {
	Name   string
	Gender string
}

type genderType struct {
	Gender string
}

func main() {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	actual := []personType{}
	lo.Filter(people, &actual, func(person interface{}) bool {
		return person.(personType).Name == "Ken"
	})
	fmt.Println(actual)
	mapThenOmitThenMap()
}

func mapThenOmitThenMap() {
	people := []personType{
		{Name: "Ken", Gender: "Male"},
		{Name: "Cythia", Gender: "Female"},
		{Name: "Kumiko", Gender: "Female"},
	}
	var actual []string
	lo.Map(people, &actual, func(p interface{}) interface{} {
		val := p.(personType)
		ret := val.Name + " " + val.Gender
		return ret
	})
	lo.Omit(actual, &actual, func(e interface{}) bool {
		val := e.(string)
		return strings.Contains(val, "K")
	})
	var actual2 []genderType
	lo.Map(actual, &actual2, func(e interface{}) interface{} {
		return genderType{Gender: e.(string)}
	})
	fmt.Println(actual2)
}
