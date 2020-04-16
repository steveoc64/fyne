package binding

import (
	"fmt"
	"reflect"
	"strconv"
)

// Int
type Int64 struct {
	Observer
	value int64
}

func NewInt(value int64) *Int64 {
	return &Int64{
		value: value,
	}
}

func (f *Int64) Kind() reflect.Kind {
	return reflect.Int64
}

func (f *Int64) Get() reflect.Value {
	if f == nil {
		return reflect.ValueOf(0)
	}
	return reflect.ValueOf(f.value)
}

func (f *Int64) Set(v reflect.Value) {
	if f == nil {
		return
	}
	value := v.Int()
	if value != f.value {
		f.value = value
		f.Update()
	}
}

func (f *Int64) GetInt64() int64 {
	if f == nil {
		return 0
	}
	return f.value
}

func (f *Int64) SetInt64(value int64) {
	if f == nil {
		return
	}
	f.value = value
	f.Update()
}

// Int64Handler wraps any existing handler and
// returns a new handler that always Gets/Sets int64s
// regardless of the next-in-chain type
func Int64Handler(h Handler) WrapHandler {
	return NewHandler(
		h,
		reflect.Int64,
		func(v reflect.Value) reflect.Value {
			if v.Kind() == reflect.Int64 {
				return v
			}
			// v is anything, return int
			str := fmt.Sprintf("%v", v)
			f, _ := strconv.ParseInt(str, 10, 64)
			return reflect.ValueOf(f)
		},
		func(v reflect.Value) reflect.Value {
			// v is int, convert it into the target type
			switch h.Kind() {
			case reflect.Int64:
				return v
			case reflect.Float64:
				return reflect.ValueOf(int64(v.Float()))
			case reflect.String:
				str := fmt.Sprintf("%v", v)
				return reflect.ValueOf(str)
			}
			return reflect.ValueOf("")
		})
}
