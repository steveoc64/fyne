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
func NewBinding(o Observable, el Notifiable, h Handler) *Binding {
	b := &Binding{o, el, h}
	o.AddListener(b)
	// kick off an update on the widget to synch it with the data
	if el != nil {
		go el.Notify(b)
	}
	return b
}
