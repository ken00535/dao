package dao

import (
	"reflect"
)

func Filter(source interface{}, out interface{}, callback func(element interface{}) bool) {
	sourceValue := reflect.ValueOf(source)
	temp := reflect.ValueOf(out).Elem()
	for i := 0; i < sourceValue.Len(); i++ {
		if callback(sourceValue.Index(i).Interface()) {
			temp = reflect.Append(temp, sourceValue.Index(i))
		}
	}
	outValue := reflect.ValueOf(out).Elem()
	outValue.Set(temp)
}
