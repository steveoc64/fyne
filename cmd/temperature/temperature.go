package main

import (
	"fyne.io/fyne/binding"
)

const KC = 273.15

// Temperature is a listenable object that stores a base temp in Kelvin
// and which provides getters and setters for Celcius and Farenheit computed values
type Temperature struct {
	binding.Float64
}

func NewTemperature() *Temperature {
	return &Temperature{}
}

func (t *Temperature) GetCelcius() float64 {
	return t.GetFloat64() - KC
}

func (t *Temperature) SetCelcius(f float64) {
	t.SetFloat64(f + KC)
}

func (t *Temperature) GetFarenheit() float64 {
	return (t.GetFloat64()-KC)*1.8 + 32.0
}

func (t *Temperature) SetFarenheit(f float64) {
	t.SetFloat64(((f - 32) / 1.8) + KC)
}
