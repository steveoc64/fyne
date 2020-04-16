package binding

import "reflect"

type WrapHandler struct {
	parent Handler
	kind   reflect.Kind
	get    func(reflect.Value) reflect.Value
	set    func(reflect.Value) reflect.Value
}

func NewHandler(parent Handler, kind reflect.Kind, get func(reflect.Value) reflect.Value, set func(reflect.Value) reflect.Value) WrapHandler {
	return WrapHandler{
		parent: parent,
		kind:   kind,
		get:    get,
		set:    set,
	}
}

func (w WrapHandler) Kind() reflect.Kind {
	return w.kind
}

func (w WrapHandler) Get() reflect.Value {
	return w.get(w.parent.Get())
}

func (w WrapHandler) Set(v reflect.Value) {
	if w.set == nil {
		// is a read only wrapper, cannot set data on this pipe
		return
	}
	w.parent.Set(w.set(v))
}
