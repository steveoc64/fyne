package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	// t only knows about Celcius and Farenheit
	// and it only deals with Float64s
	t := NewTemperature()

	// Some user-defined adapters / middleware, because we dont
	// have access to change the code in the Temperature object
	celciusAdapter := binding.AsString(Celcius(t), "%.0f")
	farenheitAdapter := binding.AsString(Farenheit(t), "%.2f")
	//kelvinAdapter := binding.AsString(Kelvinator(t), "%.4f")

	// Build the UI
	w := a.NewWindow("Celcius Farenheit Demo")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Celcius"),
		widget.NewEntry().BindText(celciusAdapter),
		widget.NewLabel("Farenheit"),
		widget.NewEntry().BindText(farenheitAdapter),
		//widget.NewLabel("Kelvin"),
		//widget.NewLabel("").BindText(kelvinAdapter),
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))

	w.ShowAndRun()
}
