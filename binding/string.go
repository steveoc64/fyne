package binding

import (
	"fmt"
	"reflect"
	"strconv"
)

// String implements an observable and a string handler
type String struct {
	Observer
	value string
}

func NewString(value string) *String {
	return &String{
		value: value,
	}
}

func (s *String) Kind() reflect.Kind {
	return reflect.String
}

// Get/Set pair to implement a Handler
func (s *String) Get() reflect.Value {
	if s == nil {
		return reflect.ValueOf("")
	}
	return reflect.ValueOf(s.value)
}

// Set/Get pair to implement a Handler
func (s *String) Set(v reflect.Value) {
	if s == nil {
		return
	}
	s.value = v.Elem().String()
	s.Update()
}

func (s *String) GetString() string {
	if s == nil {
		return ""
	}
	return s.value
}

func (s *String) SetString(value string) {
	if s == nil {
		return
	}
	s.value = value
	s.Update()
}

// StringHandler wraps any existing handler and
// returns a new handler that always Gets/Sets strings
// regardless of the next-in-chain type
func StringHandler(h Handler) WrapHandler {
	return NewHandler(
		h,
		reflect.String,
		func(v reflect.Value) reflect.Value {
			// v is anything, return string, which is easy
			str := fmt.Sprintf("%v", v)
			return reflect.ValueOf(str)
		},
		func(v reflect.Value) reflect.Value {
			// v is string, convert it into the target type
			switch h.Kind() {
			// massive set of cases .... take a string, convert to kind
			case reflect.String:
				return v
			case reflect.Float64:
				f, _ := strconv.ParseFloat(v.String(), 64)
				return reflect.ValueOf(f)
			}
			return reflect.ValueOf("")
		})
}

// Currency handler provides a filter to convert numeric data
// into currency format (2 fixed decimals)
func Currency(h Handler) WrapHandler {
	return NewHandler(
		h,
		reflect.String,
		func(v reflect.Value) reflect.Value {
			str := fmt.Sprintf("%v", v)
			f, _ := strconv.ParseFloat(str, 64)
			str = fmt.Sprintf("%.02f", f)
			return reflect.ValueOf(str)
		},
		func(v reflect.Value) reflect.Value {
			// v is string, convert it into the target type
			switch h.Kind() {
			// massive set of cases .... take a string, convert to kind
			case reflect.String:
				return v
			case reflect.Float64:
				f, _ := strconv.ParseFloat(v.String(), 64)
				return reflect.ValueOf(f)
			}
			return reflect.ValueOf("")
		})
}

// FloatString with custom format param
func FloatString(h Handler, format string) WrapHandler {
	return NewHandler(
		h,
		reflect.String,
		func(v reflect.Value) reflect.Value {
			str := fmt.Sprintf("%v", v)
			f, _ := strconv.ParseFloat(str, 64)
			str = fmt.Sprintf(format, f)
			return reflect.ValueOf(str)
		},
		func(v reflect.Value) reflect.Value {
			// v is string, convert it into the target type
			switch h.Kind() {
			// massive set of cases .... take a string, convert to kind
			case reflect.String:
				return v
			case reflect.Float64:
				f, _ := strconv.ParseFloat(v.String(), 64)
				return reflect.ValueOf(f)
			}
			return reflect.ValueOf("")
		})
}
