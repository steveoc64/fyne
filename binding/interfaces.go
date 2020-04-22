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

// A Data is anything that implements both
// Observable and Handler interfaces
// Its the thing that widgets Bind() to
type Data interface {
	Observable
	Handler
}

// A Notifiable is anything that can recieve a Notify() callback
type Notifiable interface {
	Notify(*Binding)
}
