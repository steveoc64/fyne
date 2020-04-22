package binding

import (
	"fmt"
	"reflect"

	"fyne.io/fyne"
)

// Struct implements an observable and a string handler
type Struct struct {
	Observer
	kind        reflect.Kind
	value       reflect.Value
	elementType reflect.Type
}

func NewStruct(data interface{}) *Struct {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Ptr {
		fyne.LogError("NewStruct()", fmt.Errorf("value must be a pointer, got %s", value.Type()))
		return nil
	}
	if value.Elem().Kind() != reflect.Struct {
		fyne.LogError("NewStruct()", fmt.Errorf("value must point to a struct, got ptr to %s", value.Elem().Type()))
		return nil
	}
	s := &Struct{
		kind:        reflect.Struct,
		value:       value,
		elementType: value.Elem().Type(),
	}
	return s
}

// Get/Set pair to implement a Handler
func (s *Struct) Get() reflect.Value {
	return s.value
}

// Set/Get pair to implement a Handler
func (s *Struct) Set(value reflect.Value) {
	isPtr := false
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
		isPtr = true
	}
	if value.Kind() != reflect.Struct {
		fyne.LogError("Struct.Set()", fmt.Errorf("value must be a struct of type %s, got to %s (is pointer %v)",
			s.elementType.String(),
			value.Type(),
			isPtr))
		return
	}
	if value.Type() != s.elementType {
		fyne.LogError("Struct.Set()", fmt.Errorf("value must be a struct of type %s, got to %s (is pointer %v)",
			s.elementType.String(),
			value.Type(),
			isPtr))
	}
	if !s.value.Elem().CanSet() {
		println("can set")
	}
	s.value.Elem().Set(value)
	//s.value = value
	s.Update()
}

// SetValue allows setting from a ptr to a struct
func (s *Struct) SetValue(data interface{}) {
	s.Set(reflect.ValueOf(data))
}

// Kind stamps this object of kind Struct
func (s *Struct) Kind() reflect.Kind {
	return reflect.Struct
}
