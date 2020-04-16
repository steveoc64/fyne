package binding

import "reflect"

// Observable is some data that can have listeners attached
type Observable interface {
	AddListener(*Binding)
	DeleteListener(*Binding)
	Update() // fire all the listeners
}

// Handler provides Get and Set methods and
// tracks internally what type it operates with
type Handler interface {
	Kind() reflect.Kind
	Get() reflect.Value
	Set(reflect.Value)
}

// A Bindable is anything that implements both
// Observable and Handler interfaces
type Bindable interface {
	Observable
	Handler
}

// Notifiable is an object that gets notified when data changes
// typically a widget - but could be any UI element
type Notifiable interface {
	Notify(*Binding)
}
