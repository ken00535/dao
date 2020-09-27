package slicehandler

import (
	"testing"
)

func BenchmarkFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		people := []Person{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
		actual := []Person{}
		Filter(people, &actual, func(person interface{}) bool {
			return person.(Person).Name == "Ken"
		})
	}
}

func BenchmarkFilterClassic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		people := []Person{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
		actual := []Person{}
		for j := 0; j < len(people); j++ {
			if people[j].Name == "Ken" {
				actual = append(actual, people[j])
			}
		}
	}
}
