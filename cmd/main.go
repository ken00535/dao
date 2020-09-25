package main

import "fmt"

type People struct {
	Name    string
	Address string
}

func main() {
	num2 := []People{{Name: "Ken", Address: "TW"}, {Name: "Cythia"}}
	new := []People{}
	Filter(num2, &new, func(element interface{}) bool {
		if element.(People).Name == "Ken" {
			return true
		}
		return false
	})
	fmt.Println(new)
	for i := 0; i < len(new); i++ {
		fmt.Println(new[i])
	}
}
