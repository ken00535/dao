package main

import (
	"fmt"

	dao "github.com/ken00535/dao/pkg"
)

type People struct {
	Name    string
	Address string
}

func main() {
	num2 := []People{{Name: "Ken", Address: "TW"}, {Name: "Cythia"}}
	new := []People{}
	dao.Filter(num2, &new, func(element interface{}) bool {
		return element.(People).Name == "Ken"
	})
	for i := 0; i < len(new); i++ {
		fmt.Println(new[i])
	}

	new2 := People{}
	dao.Find(num2, &new2, func(element interface{}) bool {
		return element.(People).Name == "Cythia"
	})
	fmt.Println(new2)

	new3 := []People{}
	dao.Map(num2, &new3, func(element interface{}) interface{} {
		val := element.(People)
		val.Name = val.Name + "_EX"
		return val
	})
	fmt.Println(new3)

	new4 := People{}
	dao.Reduce(num2, &new4, func(m interface{}, n interface{}) interface{} {
		people := People{}
		people.Name = m.(People).Name + n.(People).Name
		return people
	})
	fmt.Println(new4)
}
