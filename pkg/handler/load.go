package lodash

import "reflect"

// Load a data source and convert to struct that can be used.
func Load(source interface{}, out interface{}, callback func(element interface{}) bool) {
	sourceValue := reflect.ValueOf(source)
	temp := reflect.ValueOf(out).Elem()
	for i := 0; i < sourceValue.Len(); i++ {
		if callback(sourceValue.Index(i).Interface()) {
			temp.Set(sourceValue.Index(i))
			break
		}
	}
}

// End output result
func End(source interface{}, out interface{}, callback func(element interface{}) bool) {
	sourceValue := reflect.ValueOf(source)
	temp := reflect.ValueOf(out).Elem()
	for i := 0; i < sourceValue.Len(); i++ {
		if callback(sourceValue.Index(i).Interface()) {
			temp.Set(sourceValue.Index(i))
			break
		}
	}
}
