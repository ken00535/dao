package main

import (
	"fmt"

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
	lo.Start(people).Filter(func(person interface{}) bool {
		return person.(personType).Name == "Ken"
	}).End(&actual)
	fmt.Println(actual)
}
