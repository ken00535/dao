package lodash

import (
	"reflect"
)

// Partial args to the fuction
func Partial(source interface{}, out interface{}, args ...interface{}) {
	sourceValue := reflect.ValueOf(source)
	outElem := reflect.ValueOf(out).Elem()
	var argsArr []reflect.Value
	for i := 0; i < len(args); i++ {
		arg := args[i]
		argsArr = append(argsArr, reflect.ValueOf(arg))
	}

	fn := func(outArgs []reflect.Value) []reflect.Value {
		outLen := len(outArgs)
		for i := 0; i < outLen; i++ {
			argsArr = append(argsArr, outArgs[i])
		}
		return sourceValue.Call(argsArr)
	}

	outElem.Set(reflect.MakeFunc(outElem.Type(), fn))
}
