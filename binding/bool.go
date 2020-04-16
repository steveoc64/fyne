package binding

import (
	"fmt"
	"reflect"
	"strconv"
)

// Bool
type Bool struct {
	Observer
	value bool
}

func NewBool(value bool) *Bool {
	return &Bool{
		value: value,
	}
}

func (f *Bool) Kind() reflect.Kind {
	return reflect.Bool
}

func (f *Bool) Get() reflect.Value {
	if f == nil {
		return reflect.ValueOf(0)
	}
	return reflect.ValueOf(f.value)
}

func (f *Bool) Set(v reflect.Value) {
	if f == nil {
		return
	}
	value := v.Bool()
	if value != f.value {
		f.value = value
		f.Update()
	}
}

func (f *Bool) GetBool() bool {
	if f == nil {
		return false
	}
	return f.value
}

func (f *Bool) SetBool(value bool) {
	if f == nil {
		return
	}
	f.value = value
	f.Update()
}

// BoolHandler wraps any existing handler and
// returns a new handler that always Gets/Sets bools
// regardless of the next-in-chain type
func BoolHandler(h Handler) WrapHandler {
	return NewHandler(
		h,
		reflect.Bool,
		func(v reflect.Value) reflect.Value {
			if v.Kind() == reflect.Bool {
				return v
			}
			// v is anything, return bool
			str := fmt.Sprintf("%v", v)
			b, _ := strconv.ParseBool(str)
			return reflect.ValueOf(b)
		},
		func(v reflect.Value) reflect.Value {
			// v is bool, convert it into the target type
			switch h.Kind() {
			case reflect.Bool:
				return v
			case reflect.Float64:
				if v.Bool() {
					return reflect.ValueOf(1.0)
				}
				return reflect.ValueOf(0.0)
			case reflect.Int64:
				if v.Bool() {
					return reflect.ValueOf(1)
				}
				return reflect.ValueOf(0)
			case reflect.String:
				if v.Bool() {
					return reflect.ValueOf("true")
				}
				return reflect.ValueOf("false")
			}
			return reflect.ValueOf(false)
		})
}
