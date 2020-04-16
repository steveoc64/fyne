package binding

import (
	"fmt"
	"reflect"
	"strconv"
)

// Float64
type Float64 struct {
	Observer
	value float64
}

func NewFloat64(value float64) *Float64 {
	return &Float64{
		value: value,
	}
}

func (f *Float64) Kind() reflect.Kind {
	return reflect.Float64
}

func (f *Float64) Get() reflect.Value {
	if f == nil {
		return reflect.ValueOf(0)
	}
	return reflect.ValueOf(f.value)
}

func (f *Float64) Set(v reflect.Value) {
	if f == nil {
		return
	}
	value := v.Float()
	if value != f.value {
		f.value = value
		f.Update()
	}
}

func (f *Float64) GetFloat64() float64 {
	if f == nil {
		return 0.0
	}
	return f.value
}

func (f *Float64) SetFloat64(value float64) {
	if f == nil {
		return
	}
	f.value = value
	f.Update()
}

// FloatHandler wraps any existing handler and
// returns a new handler that always Gets/Sets float64s
// regardless of the next-in-chain type
func Float64Handler(h Handler) WrapHandler {
	return NewHandler(
		h,
		reflect.Float64,
		func(v reflect.Value) reflect.Value {
			if v.Kind() == reflect.Float64 {
				return v
			}
			// v is anything, return float
			str := fmt.Sprintf("%v", v)
			f, _ := strconv.ParseFloat(str, 64)
			return reflect.ValueOf(f)
		},
		func(v reflect.Value) reflect.Value {
			// v is float, convert it into the target type
			switch h.Kind() {
			case reflect.Float64:
				return v
			case reflect.Int64:
				return reflect.ValueOf(float64(v.Int()))
			case reflect.String:
				str := fmt.Sprintf("%v", v)
				return reflect.ValueOf(str)
			}
			return reflect.ValueOf("")
		})
}
