package main

import (
	"reflect"

	"fyne.io/fyne/app"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/widget"
)

// 7GUI temp converter again

const KC = 273.15

var temperature = binding.NewFloat64(0)

var celcius = binding.NewHandler(
	temperature,
	reflect.Float64,
	func(v reflect.Value) reflect.Value {
		return reflect.ValueOf(v.Float() - KC)
	},
	func(v reflect.Value) reflect.Value {
		return reflect.ValueOf(v.Float() + KC)
	},
)

var farenheit = binding.NewHandler(
	temperature,
	reflect.Float64,
	func(v reflect.Value) reflect.Value {
		return reflect.ValueOf(((v.Float() - KC) * 1.8) + 32.0)
	},
	func(v reflect.Value) reflect.Value {
		return reflect.ValueOf((v.Float()-32.0)/1.8 + KC)
	},
)

func main() {
	a := app.New()

	w := a.NewWindow("Temp Converter")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Kelvin"),
		widget.NewEntry().Bind(temperature).Handler(binding.FloatString(temperature, "%.2f")),
		widget.NewLabel("Celcius"),
		widget.NewEntry().Bind(temperature).Handler(celcius),
		widget.NewLabel("Celcius Currency"),
		widget.NewEntry().Bind(temperature).Handler(binding.Currency(celcius)),
		widget.NewLabel("Farenheit"),
		widget.NewEntry().Bind(temperature).Handler(binding.FloatString(farenheit, "%.8f")),
		widget.NewLabel("Temperature Slider"),
		widget.NewSlider(0, 10000).Bind(temperature),
	))
	w.ShowAndRun()
}
