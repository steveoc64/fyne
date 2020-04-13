package binding

// go:generate go run gen.go

import "sync"

// Binding is the base interface of the Data Binding API.
type Binding interface {
	Get() string
	Set(string)
	AddListener(Notifiable)
	DeleteListener(Notifiable)
}

// Item implements a data binding for a single item.
type Item struct {
	Binding
	sync.RWMutex
	listeners []Notifiable // TODO maybe a map[Notifiable]bool would be quicker, especially for DeleteListener?
}

// Get is a noop to implement Binding
func (b *Item) Get() string {
	return ""
}

// Set is a noop to implement Binding
func (b *Item) Set(str string) {}

// AddListenerFunction adds the given function as a listener to the binding.
// The function is wrapped in the returned NotifyFunction which can be passed to DeleteListener.
func (b *Item) AddListenerFunction(listener func(Binding)) *NotifyFunction {
	n := &NotifyFunction{
		F: listener,
	}
	b.AddListener(n)
	return n
}

// AddListener adds the given listener to the binding.
func (b *Item) AddListener(listener Notifiable) {
	b.Lock()
	defer b.Unlock()
	b.listeners = append(b.listeners, listener)
}

// DeleteListener removes the given listener from the binding.
func (b *Item) DeleteListener(listener Notifiable) {
	b.Lock()
	defer b.Unlock()
	var listeners []Notifiable
	for _, l := range b.listeners {
		if l != listener {
			listeners = append(listeners, l)
		}
	}
	b.listeners = listeners
}

func (b *Item) notify() {
	b.RLock()
	defer b.RUnlock()
	for _, l := range b.listeners {
		go l.Notify(b)
	}
}

// List implements a data binding for a list of bindings.
type List struct {
	Item
	values []Binding
}

// Length returns the length of the bound list.
func (b *List) Length() int {
	return len(b.values)
}

// Get returns the binding at the given index.
func (b *List) Get(index int) Binding {
	return b.values[index]
}

// Append adds the given binding(s) to the list.
func (b *List) Append(data ...Binding) {
	b.values = append(b.values, data...)
	b.notify()
}

// Set puts the given binding into the list at the given index.
func (b *List) Set(index int, data Binding) {
	old := b.values[index]
	if old == data {
		return
	}
	b.values[index] = data
	b.notify()
}

// Map implements a data binding for a map string to binding.
type Map struct {
	Item
	values map[string]Binding
}

// Length returns the length of the bound map.
func (b *Map) Length() int {
	return len(b.values)
}

// Get returns the binding for the given key.
func (b *Map) Get(key string) (Binding, bool) {
	v, ok := b.values[key]
	return v, ok
}

// Set puts the given binding into the map at the given key.
func (b *Map) Set(key string, data Binding) {
	old, ok := b.values[key]
	if ok && old == data {
		return
	}
	b.values[key] = data
	b.notify()
}
