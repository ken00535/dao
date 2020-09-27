package dao

import (
	"reflect"
)

// Wrap the first key that correspond condition
func Wrap(arg interface{}, source interface{}, out interface{}) {
	sourceValue := reflect.ValueOf(source)
	temp := reflect.ValueOf(out).Elem()
	argValue := reflect.ValueOf(arg)
	in := []reflect.Value{argValue}

	fn := func(args []reflect.Value) (results []reflect.Value) {
		return sourceValue.Call(in)
	}

	temp2 := reflect.MakeFunc(temp.Type(), fn)

	temp.Set(temp2)
}
