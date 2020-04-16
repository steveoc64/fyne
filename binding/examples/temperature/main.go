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
		c := v.Float() - KC
		return reflect.ValueOf(c)
	},
	func(v reflect.Value) reflect.Value {
		k := v.Float() + KC
		return reflect.ValueOf(k)
	},
)

var farenheit = binding.NewHandler(
	temperature,
	reflect.Float64,
	func(v reflect.Value) reflect.Value {
		f := ((v.Float() - KC) * 1.8) + 32.0
		return reflect.ValueOf(f)
	},
	func(v reflect.Value) reflect.Value {
		k := (v.Float()-32.0)/1.8 - 32.0
		return reflect.ValueOf(k)
	},
)

func main() {
	a := app.New()

	w := a.NewWindow("Temp Converter")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Kelvin"),
		widget.NewEntry().Bind(temperature),
		widget.NewLabel("Celcius"),
		widget.NewEntry().Bind(temperature).Handler(celcius),
		widget.NewLabel("Farenheit"),
		widget.NewEntry().Bind(temperature).Handler(farenheit),
		widget.NewLabel("Temperature Slider"),
		//widget.NewSlider(0, 10000).Bind(temperature),
	))
	w.ShowAndRun()
}
