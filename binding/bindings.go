// auto-generated
// **** THIS FILE IS AUTO-GENERATED, PLEASE DO NOT EDIT IT **** //

package binding

import (
	"fyne.io/fyne"
	"net/url"
)

type BoolBinding struct {
	ItemBinding
	value bool
}

func NewBoolBinding(value bool) *BoolBinding {
	return &BoolBinding{value: value}
}

func (b *BoolBinding) Get() bool {
	return b.value
}

func (b *BoolBinding) Set(value bool) {
	if b.value != value {
		b.value = value
		b.notify()
	}
}

func (b *BoolBinding) AddListener(listener func(bool)) {
	b.addListener(func() {
		listener(b.value)
	})
}

type Float64Binding struct {
	ItemBinding
	value float64
}

func NewFloat64Binding(value float64) *Float64Binding {
	return &Float64Binding{value: value}
}

func (b *Float64Binding) Get() float64 {
	return b.value
}

func (b *Float64Binding) Set(value float64) {
	if b.value != value {
		b.value = value
		b.notify()
	}
}

func (b *Float64Binding) AddListener(listener func(float64)) {
	b.addListener(func() {
		listener(b.value)
	})
}

type IntBinding struct {
	ItemBinding
	value int
}

func NewIntBinding(value int) *IntBinding {
	return &IntBinding{value: value}
}

func (b *IntBinding) Get() int {
	return b.value
}

func (b *IntBinding) Set(value int) {
	if b.value != value {
		b.value = value
		b.notify()
	}
}

func (b *IntBinding) AddListener(listener func(int)) {
	b.addListener(func() {
		listener(b.value)
	})
}

type Int64Binding struct {
	ItemBinding
	value int64
}

func NewInt64Binding(value int64) *Int64Binding {
	return &Int64Binding{value: value}
}

func (b *Int64Binding) Get() int64 {
	return b.value
}

func (b *Int64Binding) Set(value int64) {
	if b.value != value {
		b.value = value
		b.notify()
	}
}

func (b *Int64Binding) AddListener(listener func(int64)) {
	b.addListener(func() {
		listener(b.value)
	})
}

type ResourceBinding struct {
	ItemBinding
	value fyne.Resource
}

func NewResourceBinding(value fyne.Resource) *ResourceBinding {
	return &ResourceBinding{value: value}
}

func (b *ResourceBinding) Get() fyne.Resource {
	return b.value
}

func (b *ResourceBinding) Set(value fyne.Resource) {
	if b.value != value {
		b.value = value
		b.notify()
	}
}

func (b *ResourceBinding) AddListener(listener func(fyne.Resource)) {
	b.addListener(func() {
		listener(b.value)
	})
}

type RuneBinding struct {
	ItemBinding
	value rune
}

func NewRuneBinding(value rune) *RuneBinding {
	return &RuneBinding{value: value}
}

func (b *RuneBinding) Get() rune {
	return b.value
}

func (b *RuneBinding) Set(value rune) {
	if b.value != value {
		b.value = value
		b.notify()
	}
}

func (b *RuneBinding) AddListener(listener func(rune)) {
	b.addListener(func() {
		listener(b.value)
	})
}

type StringBinding struct {
	ItemBinding
	value string
}

func NewStringBinding(value string) *StringBinding {
	return &StringBinding{value: value}
}

func (b *StringBinding) Get() string {
	return b.value
}

func (b *StringBinding) Set(value string) {
	if b.value != value {
		b.value = value
		b.notify()
	}
}

func (b *StringBinding) AddListener(listener func(string)) {
	b.addListener(func() {
		listener(b.value)
	})
}

type URLBinding struct {
	ItemBinding
	value *url.URL
}

func NewURLBinding(value *url.URL) *URLBinding {
	return &URLBinding{value: value}
}

func (b *URLBinding) Get() *url.URL {
	return b.value
}

func (b *URLBinding) Set(value *url.URL) {
	if b.value != value {
		b.value = value
		b.notify()
	}
}

func (b *URLBinding) AddListener(listener func(*url.URL)) {
	b.addListener(func() {
		listener(b.value)
	})
}
