package dao

import (
	"reflect"
)

// Find the first element that correspond condition
func Find(source interface{}, out interface{}, callback func(element interface{}) bool) {
	sourceValue := reflect.ValueOf(source)
	temp := reflect.ValueOf(out).Elem()
	for i := 0; i < sourceValue.Len(); i++ {
		if callback(sourceValue.Index(i).Interface()) {
			temp.Set(sourceValue.Index(i))
			break
		}
	}
}

// Filter return the elements that correspond condition
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

// ForEach visit each element and call callback
func ForEach(source interface{}, out interface{}, callback func(element interface{})) {
	sourceValue := reflect.ValueOf(source)
	for i := 0; i < sourceValue.Len(); i++ {
		callback(sourceValue.Index(i).Interface())
	}
}

// Map element to a new slice
func Map(source interface{}, out interface{}, callback func(element interface{}) interface{}) {
	sourceValue := reflect.ValueOf(source)
	temp := reflect.ValueOf(out).Elem()
	for i := 0; i < sourceValue.Len(); i++ {
		mapVal := callback(sourceValue.Index(i).Interface())
		temp = reflect.Append(temp, reflect.ValueOf(mapVal))
	}
	outValue := reflect.ValueOf(out).Elem()
	outValue.Set(temp)
}

// Reduce slice to a value
func Reduce(source interface{}, out interface{}, callback func(interface{}, interface{}) interface{}) {
	sourceValue := reflect.ValueOf(source)
	reduce := callback(sourceValue.Index(0).Interface(), sourceValue.Index(1).Interface())
	for i := 2; i < sourceValue.Len(); i++ {
		reduce = callback(reduce, sourceValue.Index(i).Interface())
	}
	reduceValue := reflect.ValueOf(reduce)
	outValue := reflect.ValueOf(out).Elem()
	outValue.Set(reduceValue)
}

// Omit delete the elements that correspond condition
func Omit(source interface{}, out interface{}, callback func(element interface{}) bool) {
	sourceValue := reflect.ValueOf(source)
	temp := reflect.ValueOf(out).Elem()
	for i := 0; i < sourceValue.Len(); i++ {
		if !callback(sourceValue.Index(i).Interface()) {
			temp = reflect.Append(temp, sourceValue.Index(i))
		}
	}
	outValue := reflect.ValueOf(out).Elem()
	outValue.Set(temp)
}
