// auto-generated
// **** THIS FILE IS AUTO-GENERATED, PLEASE DO NOT EDIT IT **** //

package binding

import (
	"math"
	"net/url"

	"fyne.io/fyne"
)

// Bool defines a data binding for a bool
type IBool interface {
	Binding
	GetBool() bool
	SetBool(bool)
	AddBoolListener(func(bool)) *NotifyFunction
}

// Bool implements an IBool interface
type Bool struct {
	Item

	value bool
}

// NewBool creates a new binding with the given value.
func NewBool(value bool) IBool {
	return &Bool{value: value}
}

// Get returns the bound value.
func (b *Bool) GetBool() bool {
	return b.value
}

// Set updates the bound value.
func (b *Bool) SetBool(value bool) {
	if b.value != value {
		b.value = value
		b.notify()
	}
}

// AddBoolListener adds the given function as a listener to the binding.
// The function is wrapped in the returned NotifyFunction which can be passed to DeleteListener.
func (b *Bool) AddBoolListener(listener func(bool)) *NotifyFunction {
	return b.AddListenerFunction(func(Binding) {
		listener(b.value)
	})
}

// Float64 defines a data binding for a float64
type IFloat64 interface {
	Binding
	GetFloat64() float64
	SetFloat64(float64)
	AddFloat64Listener(func(float64)) *NotifyFunction
}

// Float64 implements an IFloat64 interface
type Float64 struct {
	Item
	value float64
}

// NewFloat64 creates a new binding with the given value.
func NewFloat64(value float64) IFloat64 {
	return &Float64{value: value}
}

// Get returns the bound value.
func (b *Float64) GetFloat64() float64 {
	return b.value
}

// Set updates the bound value.
func (b *Float64) SetFloat64(value float64) {
	if math.Abs(b.value-value) > 0.0001 {
		b.value = value
		b.notify()
	}
}

// AddFloat64Listener adds the given function as a listener to the binding.
// The function is wrapped in the returned NotifyFunction which can be passed to DeleteListener.
func (b *Float64) AddFloat64Listener(listener func(float64)) *NotifyFunction {
	return b.AddListenerFunction(func(Binding) {
		listener(b.value)
	})
}

// Int defines a data binding for a int
type IInt interface {
	Binding
	GetInt() int
	SetInt(int)
	AddIntListener(func(int)) *NotifyFunction
}

// Int implements an IInt interface
type Int struct {
	Item
	value int
}

// NewInt creates a new binding with the given value.
func NewInt(value int) IInt {
	return &Int{value: value}
}

// Get returns the bound value.
func (b *Int) GetInt() int {
	return b.value
}

// Set updates the bound value.
func (b *Int) SetInt(value int) {
	if b.value != value {
		b.value = value
		b.notify()
	}
}

// AddIntListener adds the given function as a listener to the binding.
// The function is wrapped in the returned NotifyFunction which can be passed to DeleteListener.
func (b *Int) AddIntListener(listener func(int)) *NotifyFunction {
	return b.AddListenerFunction(func(Binding) {
		listener(b.value)
	})
}

// Int64 defines a data binding for a int64
type IInt64 interface {
	Binding
	GetInt64() int64
	SetInt64(int64)
	AddInt64Listener(func(int64)) *NotifyFunction
}

// Int64 implements an IInt64 interface
type Int64 struct {
	Item
	value int64
}

// NewInt64 creates a new binding with the given value.
func NewInt64(value int64) IInt64 {
	return &Int64{value: value}
}

// Get returns the bound value.
func (b *Int64) GetInt64() int64 {
	return b.value
}

// Set updates the bound value.
func (b *Int64) SetInt64(value int64) {
	if b.value != value {
		b.value = value
		b.notify()
	}
}

// AddInt64Listener adds the given function as a listener to the binding.
// The function is wrapped in the returned NotifyFunction which can be passed to DeleteListener.
func (b *Int64) AddInt64Listener(listener func(int64)) *NotifyFunction {
	return b.AddListenerFunction(func(Binding) {
		listener(b.value)
	})
}

// Resource defines a data binding for a fyne.Resource
type IResource interface {
	Binding
	GetResource() fyne.Resource
	SetResource(fyne.Resource)
	AddResourceListener(func(fyne.Resource)) *NotifyFunction
}

// Resource implements an IResource interface
type Resource struct {
	Item
	value fyne.Resource
}

// NewResource creates a new binding with the given value.
func NewResource(value fyne.Resource) IResource {
	return &Resource{value: value}
}

// Get returns the bound value.
func (b *Resource) GetResource() fyne.Resource {
	return b.value
}

// Set updates the bound value.
func (b *Resource) SetResource(value fyne.Resource) {
	if b.value != value {
		b.value = value
		b.notify()
	}
}

// AddResourceListener adds the given function as a listener to the binding.
// The function is wrapped in the returned NotifyFunction which can be passed to DeleteListener.
func (b *Resource) AddResourceListener(listener func(fyne.Resource)) *NotifyFunction {
	return b.AddListenerFunction(func(Binding) {
		listener(b.value)
	})
}

// Rune defines a data binding for a rune
type IRune interface {
	Binding
	GetRune() rune
	SetRune(rune)
	AddRuneListener(func(rune)) *NotifyFunction
}

// Rune implements an IRune interface
type Rune struct {
	Item
	value rune
}

// NewRune creates a new binding with the given value.
func NewRune(value rune) IRune {
	return &Rune{value: value}
}

// Get returns the bound value.
func (b *Rune) GetRune() rune {
	return b.value
}

// Set updates the bound value.
func (b *Rune) SetRune(value rune) {
	if b.value != value {
		b.value = value
		b.notify()
	}
}

// AddRuneListener adds the given function as a listener to the binding.
// The function is wrapped in the returned NotifyFunction which can be passed to DeleteListener.
func (b *Rune) AddRuneListener(listener func(rune)) *NotifyFunction {
	return b.AddListenerFunction(func(Binding) {
		listener(b.value)
	})
}

// String defines a data binding for a string
type IString interface {
	Binding
	GetString() string
	SetString(string)
	AddStringListener(func(string)) *NotifyFunction
}

// String implements an IString interface
type String struct {
	Item
	value string
}

// NewString creates a new binding with the given value.
func NewString(value string) IString {
	return &String{value: value}
}

// Get returns the bound value.
func (b *String) GetString() string {
	return b.value
}

// Set updates the bound value.
func (b *String) SetString(value string) {
	if b.value != value {
		b.value = value
		b.notify()
	}
}

// AddStringListener adds the given function as a listener to the binding.
// The function is wrapped in the returned NotifyFunction which can be passed to DeleteListener.
func (b *String) AddStringListener(listener func(string)) *NotifyFunction {
	return b.AddListenerFunction(func(Binding) {
		listener(b.value)
	})
}

// URL defines a data binding for a *url.URL
type IURL interface {
	Binding
	GetURL() *url.URL
	SetURL(*url.URL)
	AddURLListener(func(*url.URL)) *NotifyFunction
}

// URL implements an IURL interface
type URL struct {
	Item
	value *url.URL
}

// NewURL creates a new binding with the given value.
func NewURL(value *url.URL) IURL {
	return &URL{value: value}
}

// Get returns the bound value.
func (b *URL) GetURL() *url.URL {
	return b.value
}

// Set updates the bound value.
func (b *URL) SetURL(value *url.URL) {
	if b.value != value {
		b.value = value
		b.notify()
	}
}

// AddURLListener adds the given function as a listener to the binding.
// The function is wrapped in the returned NotifyFunction which can be passed to DeleteListener.
func (b *URL) AddURLListener(listener func(*url.URL)) *NotifyFunction {
	return b.AddListenerFunction(func(Binding) {
		listener(b.value)
	})
}
