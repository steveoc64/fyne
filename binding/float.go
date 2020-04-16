package binding

import "reflect"

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
	f.value = v.Float()
	f.Update()
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
