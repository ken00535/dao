package main

import (
	"fmt"

	dao "github.com/ken00535/dao/pkg/handler"
)

type personType struct {
	Name   string
	Gender string
}

func main() {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	actual := []personType{}
	dao.Filter(people, &actual, func(person interface{}) bool {
		return person.(personType).Name == "Ken"
	})
	fmt.Println(actual)
}
