package main

import (
	"reflect"

	"fyne.io/fyne/app"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/widget"
)

// 7GUI temp converter again

const KC = 273.15

func main() {
	a := app.New()

	temperature := binding.NewFloat64(0.0)

	celcius := binding.NewHandler(
		temperature,
		reflect.Float64,
		func(v reflect.Value) reflect.Value {
			return reflect.ValueOf(v.Float() - KC)
		},
		func(v reflect.Value) reflect.Value {
			return reflect.ValueOf(v.Float() + KC)
		},
	)

	farenheit := binding.NewHandler(
		temperature,
		reflect.Float64,
		func(v reflect.Value) reflect.Value {
			return reflect.ValueOf(((v.Float() - KC) * 1.8) + 32.0)
		},
		func(v reflect.Value) reflect.Value {
			return reflect.ValueOf((v.Float()-32.0)/1.8 + KC)
		},
	)
	w := a.NewWindow("Temp Converter")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Kelvin Raw"),
		widget.NewEntry().Bind(temperature),
		widget.NewLabel("Kelvin 2 Decimals"),
		widget.NewEntry().
			Bind(temperature).
			Handler(binding.FloatString(temperature, "%.2f")),
		widget.NewLabel("Celcius Raw"),
		widget.NewEntry().
			Bind(temperature).
			Handler(celcius),
		widget.NewLabel("Celcius 2 Decimals"),
		widget.NewEntry().
			Bind(temperature).
			Handler(binding.FloatString(temperature, "%.02f")),
		widget.NewLabel("Cost in US Dollars"),
		widget.NewEntry().
			Bind(temperature).
			Handler(binding.Currency(celcius, "$", 1.0)),
		widget.NewLabel("Cost in AUD$"),
		widget.NewEntry().
			Bind(temperature).
			Handler(binding.Currency(celcius, "AUD $ ", 0.64)),
		widget.NewLabel("Farenheit 4 Full decimals"),
		widget.NewEntry().
			Bind(temperature).
			Handler(binding.FloatString(farenheit, "%.04f")),
		widget.NewLabel("Temperature range 0-10,000k"),
		widget.NewSlider(0, 10000).Bind(temperature),
	))
	w.ShowAndRun()
}
