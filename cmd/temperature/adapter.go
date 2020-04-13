package main

import "fyne.io/fyne/binding"

// Celcius is an adapter function that takes a Temperature object
// (which is conforms to a binding.Float64)
// and wraps it to adapt to binding.String() over the Celcius computed prop
func Celcius(t *Temperature) binding.GetterSetter {
	return binding.GetterSetter{
		t,
		func() interface{} { return t.GetCelcius() },
		func(v interface{}) {
			if f, ok := v.(float64); ok {
				t.SetCelcius(f)
			}
		},
	}
}

// Celcius is an adapter function that takes a Temperature object
// (which is conforms to a binding.Float64)
// and wraps it to adapt to binding.String() over the Farenheit computed prop
func Farenheit(t *Temperature) binding.GetterSetter {
	return binding.GetterSetter{
		t,
		func() interface{} { return t.GetFarenheit() },
		func(v interface{}) {
			if f, ok := v.(float64); ok {
				t.SetFarenheit(f)
			}
		},
	}
}

// Kelvinator is a read only string adapter for the Kelvin value
func Kelvinator(t *Temperature) binding.GetterSetter {
	return binding.GetterSetter{
		t,
		func() interface{} { return t.GetFloat64() },
		func(v interface{}) {},
	}
}
