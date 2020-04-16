package binding

// Binding records an instance of a connection bewteen
// some data (type Observable)
// some widget/UI element (type Notifiable)
// the accessor methods used (type Handler)
type Binding struct {
	Data    Observable
	Element Notifiable
	Handler Handler
}

// NewBinding convenience func
func NewBinding(o Observable, e Notifiable, h Handler) *Binding {
	b := &Binding{o, e, h}
	o.AddListener(b)
	return b
}
