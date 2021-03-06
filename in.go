package goin

import (
	"fmt"
	"reflect"
)

// NewValue returns a new Value set data to x.
func NewValue(x interface{}) *Value {
	return &Value{
		data: x,
	}
}

// A Value holds a data to be checked in array, slice or map.
type Value struct {
	data interface{}
}

// In reports whether Value can be found in array/slice arr.
func (i *Value) In(arr interface{}) (bool, error) {
	s := reflect.ValueOf(arr)
	kind := s.Kind()
	len := s.Len()

	if kind != reflect.Slice && kind != reflect.Array {
		return false, fmt.Errorf(`invalid kind "%s". support only "slice" and "array"`, s.Kind())
	}

	switch reflect.TypeOf(i.data).Kind() {
	case reflect.Int:
		data := i.data.(int)
		for index := 0; index < len; index++ {
			if d, ok := s.Index(index).Interface().(int); ok {
				if data == d {
					return true, nil
				}
			}
		}
	case reflect.Int32:
		data := i.data.(int32)
		for index := 0; index < len; index++ {
			if d, ok := s.Index(index).Interface().(int32); ok {
				if data == d {
					return true, nil
				}
			}
		}
	case reflect.Int64:
		data := i.data.(int64)
		for index := 0; index < len; index++ {
			v := s.Index(index)
			if v.Kind() == reflect.Int64 {
				if s.Index(index).Int() == data {
					return true, nil
				}
			}
		}
	case reflect.Float32:
		data := i.data.(float32)
		for index := 0; index < len; index++ {
			if d, ok := s.Index(index).Interface().(float32); ok {
				if data == d {
					return true, nil
				}
			}
		}
	case reflect.Float64:
		data := i.data.(float64)
		for index := 0; index < len; index++ {
			if d, ok := s.Index(index).Interface().(float64); ok {
				if data == d {
					return true, nil
				}
			}
		}
	case reflect.String:
		data := i.data.(string)
		for index := 0; index < len; index++ {
			if d, ok := s.Index(index).Interface().(string); ok {
				if data == d {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

// InKey reports whether Value can be found as key of map arr.
func (i *Value) InKey(arr interface{}) (bool, error) {
	s := reflect.ValueOf(arr)

	if s.Kind() != reflect.Map {
		return false, fmt.Errorf(`invalid kind "%s". support only "Map"`, s.Kind())
	}

	value := reflect.ValueOf(arr).MapIndex(reflect.ValueOf(i.data))
	return value.IsValid(), nil
}
