package main

import (
	"reflect"

	"fyne.io/fyne/app"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/widget"
)

// 7GUI temp converter again

const KC = 273.15

// Temperature is a reusable data item that tracks temperature
// in kelvin, and provides custom Celcius and Farenheit handlers
type Temperature struct {
	*binding.Float64
	Celcius   binding.Handler
	Farenheit binding.Handler
}

func NewTemperature(value float64) *Temperature {
	t := &Temperature{
		Float64: binding.NewFloat64(value),
	}

	t.Celcius = binding.NewHandler(t, reflect.Float64,
		func(v reflect.Value) reflect.Value {
			return reflect.ValueOf(v.Float() - KC)
		},
		func(v reflect.Value) reflect.Value {
			return reflect.ValueOf(v.Float() + KC)
		},
	)
	t.Farenheit = binding.NewHandler(t, reflect.Float64,
		func(v reflect.Value) reflect.Value {
			return reflect.ValueOf(((v.Float() - KC) * 1.8) + 32.0)
		},
		func(v reflect.Value) reflect.Value {
			return reflect.ValueOf((v.Float()-32.0)/1.8 + KC)
		},
	)
	return t
}

func main() {
	a := app.New()

	t := NewTemperature(0.0)

	w := a.NewWindow("Mega Temp Converter")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Absolute Temp as Int"),
		widget.NewEntry().
			Bind(t).
			Handler(binding.Numberf("%d", t)).
			OnChanged(func(str string) {
				println("edit changed to", str)
			}),
		widget.NewLabel("Kelvin Raw Value"),
		widget.NewEntry().Bind(t),
		widget.NewLabel("Kelvin 2 Decimals"),
		widget.NewEntry().
			Bind(t).
			Handler(binding.Numberf("%.2f", t)),
		widget.NewLabel("Celcius Raw"),
		widget.NewEntry().
			Bind(t).
			Handler(t.Celcius),
		widget.NewLabel("Celcius 2 Decimals"),
		widget.NewEntry().
			Bind(t).
			Handler(binding.Numberf("%.2f", t.Celcius)),
		widget.NewLabel("Farenheit 4 Full decimals"),
		widget.NewEntry().
			Bind(t).
			Handler(binding.Numberf("%.04f", t.Farenheit)),
		widget.NewLabel("Temperature range 0-10,000k"),
		widget.NewSlider(0, 10000).Bind(t),
	))
	w.Show()

	wc := a.NewWindow("Temp Cost Calculator")
	wc.SetContent(widget.NewVBox(
		widget.NewLabel("Cost in US Dollars"),
		widget.NewEntry().
			Bind(t).
			Handler(binding.Currency(t.Celcius, "$", 1.0)),
		widget.NewLabel("Cost in AUD$"),
		widget.NewEntry().
			Bind(t).
			Handler(binding.Currency(t.Celcius, "AUD $ ", 0.64)),
	))
	wc.ShowAndRun()
}
