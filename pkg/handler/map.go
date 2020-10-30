package lodash

import "reflect"

// FindKey the first key that correspond condition
func (m *Type) FindKey(callback func(element interface{}) bool) *Type {
	sourceValue := m.mapElement
	iter := sourceValue.MapRange()
	tmpValue := reflect.New(sourceValue.Type().Key()).Elem()
	for iter.Next() {
		if callback(iter.Value().Interface()) {
			tmpValue.Set(iter.Key())
			break
		}
	}
	m.mapElement = tmpValue
	return m
}

// Keys get keys of map
func (m *Type) Keys() *Type {
	sourceValue := m.mapElement
	iter := sourceValue.MapRange()
	tmpValue := reflect.MakeSlice(reflect.SliceOf(sourceValue.Type().Key()), 0, 0)
	for iter.Next() {
		tmpValue = reflect.Append(tmpValue, iter.Key())
	}
	m.mapElement = tmpValue
	return m
}

// Values get values of map
func (m *Type) Values() *Type {
	sourceValue := m.mapElement
	iter := sourceValue.MapRange()
	tmpValue := reflect.MakeSlice(reflect.SliceOf(sourceValue.Type().Elem()), 0, 0)
	for iter.Next() {
		tmpValue = reflect.Append(tmpValue, iter.Value())
	}
	m.mapElement = tmpValue
	return m
}

// Pick the elemets that corresponding confition
func (m *Type) Pick(callback func(element interface{}) bool) *Type {
	sourceValue := m.mapElement
	iter := sourceValue.MapRange()
	tmpValue := reflect.MakeMap(reflect.MapOf(sourceValue.Type().Key(), sourceValue.Type().Elem()))
	for iter.Next() {
		if callback(iter.Value().Interface()) {
			tmpValue.SetMapIndex(iter.Key(), iter.Value())
		}
	}
	m.mapElement = tmpValue
	return m
}

// Merge all maps to a map
func (m *Type) Merge(other interface{}) *Type {
	sourceValue := m.mapElement
	otherValue := reflect.ValueOf(other)
	tmpValue := reflect.MakeMap(reflect.MapOf(sourceValue.Type().Key(), sourceValue.Type().Elem()))
	iter := sourceValue.MapRange()
	for iter.Next() {
		tmpValue.SetMapIndex(iter.Key(), iter.Value())
	}
	iter = otherValue.MapRange()
	for iter.Next() {
		tmpValue.SetMapIndex(iter.Key(), iter.Value())
	}
	m.mapElement = tmpValue
	return m
}
