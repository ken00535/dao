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

// Keys get keys of map
func Keys(source interface{}, out interface{}) {
	sourceValue := reflect.ValueOf(source)
	temp := reflect.ValueOf(out).Elem()
	outElem := temp
	keys := sourceValue.MapKeys()
	for i := 0; i < len(keys); i++ {
		temp = reflect.Append(temp, keys[i])
	}
	outElem.Set(temp)
}

// Pick the elemets that corresponding confition
func Pick(source interface{}, out interface{}, callback func(element interface{}) bool) {
	sourceValue := reflect.ValueOf(source)
	temp := reflect.ValueOf(out).Elem()
	outElem := temp
	iter := sourceValue.MapRange()
	for iter.Next() {
		if callback(iter.Value().Interface()) {
			temp.SetMapIndex(iter.Key(), iter.Value())
		}
	}
	outElem.Set(temp)
}
