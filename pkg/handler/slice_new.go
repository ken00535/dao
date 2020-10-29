package lodash

import (
	"reflect"
)

// Type is type
type Type struct {
	mapElememt   map[interface{}]interface{}
	sliceElement reflect.Value
}

// Start start
func Start(source interface{}) *Type {
	m := &Type{}
	m.sliceElement = reflect.ValueOf(source)
	return m
}

// Filter return the elements that correspond condition
func (m *Type) Filter(callback func(element interface{}) bool) *Type {
	sourceValue := m.sliceElement
	sliceValue := reflect.MakeSlice(sourceValue.Type(), 0, 0)
	for i := 0; i < sourceValue.Len(); i++ {
		if callback(sourceValue.Index(i).Interface()) {
			sliceValue = reflect.Append(sliceValue, sourceValue.Index(i))
		}
	}
	m.sliceElement = sliceValue
	return m
}

// End return the elements that correspond condition
func (m *Type) End(out interface{}) {
	outValue := reflect.ValueOf(out).Elem()
	outValue.Set(m.sliceElement)
}
