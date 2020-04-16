package binding

import "sync"

// Observer implements a basic observable
type Observer struct {
	sync.RWMutex // because locking
	listeners    []*Binding
}

// AddListener adds the binding
func (o *Observer) AddListener(b *Binding) {
	o.Lock()
	defer o.Unlock()
	for _, v := range o.listeners {
		if v == b {
			return // dont double up
		}
	}
	o.listeners = append(o.listeners, b)
}

// DeleteListener deletes the matching binding if found
func (o *Observer) DeleteListener(b *Binding) {
	o.Lock()
	defer o.Unlock()
	for k, v := range o.listeners {
		if v == b {
			o.listeners = append(o.listeners[:k], o.listeners[k+1:]...)
			return
		}
	}
}

// Update fires all the listeners
func (o *Observer) Update() {
	o.RLock()
	defer o.RUnlock()
	for _, v := range o.listeners {
		updateChan <- v
	}
}
