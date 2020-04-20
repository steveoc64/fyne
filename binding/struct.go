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
	value = value.Elem()
	if value.Kind() != reflect.Struct {
		fyne.LogError("NewStruct()", fmt.Errorf("value must point to a struct, got %s", value.Type()))
		return nil
	}
	s := &Struct{
		kind:        reflect.Struct,
		value:       reflect.ValueOf(data),
		elementType: value.Type(),
	}
	println("setting struct of type", s.elementType.Name())
	fmt.Printf("passed %T %v %s\n", value, value, value.Kind())
	return s
}

// Get/Set pair to implement a Handler
func (s *Struct) Get() reflect.Value {
	return s.value
}

// Set/Get pair to implement a Handler
func (s *Struct) Set(value reflect.Value) {
	// v must be a ptr or value of struct of type v.elementType
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Type() != s.elementType {
		fyne.LogError("Struct.Set()", fmt.Errorf("requires a struct of type %s, got %s", s.elementType.Name(), value.Type()))
	}
	s.Update()
}

func (s *Struct) Kind() reflect.Kind {
	return reflect.Struct
}
