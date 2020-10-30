package lodash

import "reflect"

// Type is type
type Type struct {
	sliceElement reflect.Value
	mapElement   reflect.Value
}

// Start a chain
func Start(source interface{}) *Type {
	m := &Type{}
	if reflect.TypeOf(source).Kind() == reflect.Slice {
		m.sliceElement = reflect.ValueOf(source)
	} else if reflect.TypeOf(source).Kind() == reflect.Map {
		m.mapElement = reflect.ValueOf(source)
	}
	return m
}

// End a chain and return values
func (m *Type) End(out interface{}) {
	outValue := reflect.ValueOf(out).Elem()
	if m.sliceElement.IsValid() {
		outValue.Set(m.sliceElement)
	} else if m.mapElement.IsValid() {
		outValue.Set(m.mapElement)
	}
}
