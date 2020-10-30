package lodash

import (
	"reflect"
)

// ForEach visit each element and call callback
func (m *Type) ForEach(callback func(element interface{}) bool) *Type {
	sourceValue := m.sliceElement
	for i := 0; i < sourceValue.Len(); i++ {
		callback(sourceValue.Index(i).Interface())
	}
	return m
}

// Find the first element that correspond condition
func (m *Type) Find(callback func(element interface{}) bool) *Type {
	sourceValue := m.sliceElement
	tmpValue := reflect.New(sourceValue.Index(1).Type()).Elem()
	foreachCallback := func(element interface{}) bool {
		found := callback(element)
		if found {
			tmpValue.Set(reflect.ValueOf(element))
		}
		return !found
	}
	m.ForEach(foreachCallback)
	m.sliceElement = tmpValue
	return m
}

// Filter return the elements that correspond condition
func (m *Type) Filter(callback func(element interface{}) bool) *Type {
	sourceValue := m.sliceElement
	sliceValue := reflect.MakeSlice(sourceValue.Type(), 0, 0)
	foreachCallback := func(element interface{}) bool {
		found := callback(element)
		if found {
			sliceValue = reflect.Append(sliceValue, reflect.ValueOf(element))
		}
		return true
	}
	m.ForEach(foreachCallback)
	m.sliceElement = sliceValue
	return m
}

// Omit delete the elements that correspond condition
func (m *Type) Omit(callback func(element interface{}) bool) *Type {
	sourceValue := m.sliceElement
	sliceValue := reflect.MakeSlice(sourceValue.Type(), 0, 0)
	foreachCallback := func(element interface{}) bool {
		omit := callback(element)
		if !omit {
			sliceValue = reflect.Append(sliceValue, reflect.ValueOf(element))
		}
		return true
	}
	m.ForEach(foreachCallback)
	m.sliceElement = sliceValue
	return m
}

// Map element to a new slice
func (m *Type) Map(callback func(element interface{}) interface{}) *Type {
	sourceValue := m.sliceElement
	sliceValue := reflect.MakeSlice(sourceValue.Type(), 0, 0)
	foreachCallback := func(element interface{}) bool {
		mapVal := callback(element)
		sliceValue = reflect.Append(sliceValue, reflect.ValueOf(mapVal))
		return true
	}
	m.ForEach(foreachCallback)
	m.sliceElement = sliceValue
	return m
}

// Reduce slice to a value
func (m *Type) Reduce(callback func(interface{}, interface{}) interface{}) *Type {
	sourceValue := m.sliceElement
	reduce := callback(sourceValue.Index(0).Interface(), sourceValue.Index(1).Interface())
	for i := 2; i < sourceValue.Len(); i++ {
		reduce = callback(reduce, sourceValue.Index(i).Interface())
	}
	m.sliceElement = reflect.ValueOf(reduce)
	return m
}

// Uniq remove the duplicated elements
func (m *Type) Uniq() *Type {
	record := make(map[interface{}]bool)
	sourceValue := m.sliceElement
	sliceValue := reflect.MakeSlice(sourceValue.Type(), 0, 0)
	foreachCallback := func(element interface{}) bool {
		v := element
		if _, exist := record[element]; !exist {
			record[v] = true
			sliceValue = reflect.Append(sliceValue, reflect.ValueOf(element))
		}
		return true
	}
	m.ForEach(foreachCallback)
	m.sliceElement = sliceValue
	return m
}

// Difference return slice values not included in the other given
func (m *Type) Difference(values interface{}) *Type {
	record := make(map[interface{}]bool)
	sourceValue := m.sliceElement
	sliceValue := reflect.MakeSlice(sourceValue.Type(), 0, 0)
	foreachRecord := func(element interface{}) bool {
		if _, exist := record[reflect.ValueOf(element).Interface()]; !exist {
			record[reflect.ValueOf(element).Interface()] = true
		}
		return true
	}
	foreachAppend := func(element interface{}) bool {
		if _, exist := record[element]; !exist {
			sliceValue = reflect.Append(sliceValue, reflect.ValueOf(element))
		}
		return true
	}
	Start(values).ForEach(foreachRecord)
	m.ForEach(foreachAppend)
	m.sliceElement = sliceValue
	return m
}
