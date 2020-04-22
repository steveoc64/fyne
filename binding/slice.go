package binding

import (
	"reflect"
)

// Slice implements an observable and an Int handler to get the length
// binding.Slice is basically a wrapper around binding.Observer and reflect.Slice
type Slice struct {
	Observer
	kind        reflect.Kind
	values      reflect.Value
	elementType reflect.Type
}

// NewSlice returns a new Slice of the type given by example
// eg:
//    NewSlice("") is a Slice of string
//    NewSlice(0)  is a Slice of int
//    NewSlice(&MyThing{})  is a Slice of pointers to struct MyThing
func NewSlice(example interface{}, len int, cap int) *Slice {
	t := reflect.TypeOf(example)
	return &Slice{
		kind:        reflect.Slice,
		values:      reflect.MakeSlice(reflect.SliceOf(t), len, cap),
		elementType: t,
	}
}

// Kind returns the reflection kind
func (s *Slice) Kind() reflect.Kind {
	return reflect.Slice
}

// ElementType returns the type of each element that the slice contains
func (s *Slice) ElementType() reflect.Type {
	if s == nil {
		return nil
	}
	return s.elementType
}

// Get on a slice returns its length
func (s *Slice) Get() reflect.Value {
	if s == nil {
		return reflect.ValueOf(0)
	}
	return reflect.ValueOf(s.values.Len())
}

// Set on a slice is a no-op but needs to be present to fullfil the
// Handler interface
func (s *Slice) Set(v reflect.Value) {
}

// Len returns the number of elements in the slice
func (s *Slice) Len() int {
	if s == nil {
		return 0
	}
	return s.values.Len()
}

// Index returns the element at index i
func (s *Slice) Index(i int) interface{} {
	if s == nil {
		return reflect.Zero(s.elementType)
	}
	if i < 0 || i >= s.values.Len() {
		return reflect.Zero(s.elementType)
	}
	return s.values.Index(i).Interface()
}

// Append adds elements to the slice
// Will panic if the values passed do not match the original type
func (s *Slice) Append(values ...interface{}) {
	if s == nil {
		return
	}
	for _, v := range values {
		s.values = reflect.Append(s.values, reflect.ValueOf(v))
	}
	s.Update()
}
