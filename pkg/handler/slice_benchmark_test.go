package lodash

import (
	"testing"
)

func BenchmarkFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
		actual := []personType{}
		Start(people).Filter(func(person interface{}) bool {
			return person.(personType).Name == "Ken"
		}).End(&actual)
	}
}

func BenchmarkFilterClassic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
		actual := []personType{}
		for j := 0; j < len(people); j++ {
			if people[j].Name == "Ken" {
				actual = append(actual, people[j])
			}
		}
	}
}
