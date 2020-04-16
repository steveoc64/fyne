package binding

// primitive processing loop using a queue
var updateChan = make(chan *Binding, 24000)

func init() {
	go processingLoop()
}

// run me in a go-routine to do all the notification processing please
func processingLoop() {
	for {
		select {
		case b := <-updateChan:
			b.Element.Notify(b)
		}
	}
}
