package dao

import (
	"reflect"
)

// FindKey the first key that correspond condition
func FindKey(source interface{}, out interface{}, callback func(element interface{}) bool) {
	sourceValue := reflect.ValueOf(source)
	temp := reflect.ValueOf(out).Elem()
	iter := sourceValue.MapRange()
	for iter.Next() {
		if callback(iter.Value().Interface()) {
			temp.Set(iter.Key())
			break
		}
	}
}
