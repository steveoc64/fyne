package binding

// primitive processing loop using a queue
var updateChan = make(chan *Binding, 24000)

func init() {
	go func() {
		for {
			select {
			case b := <-updateChan:
				b.Element.Notify(b)
			}
		}
	}()
}
