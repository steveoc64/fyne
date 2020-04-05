package widget

import (
	"testing"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/test"
	"github.com/stretchr/testify/assert"
)

func TestProgressBar_SetValue(t *testing.T) {
	bar := NewProgressBar()

	assert.Equal(t, 0.0, bar.Min)
	assert.Equal(t, 1.0, bar.Max)
	assert.Equal(t, 0.0, bar.Value)

	bar.SetValue(.5)
	assert.Equal(t, .5, bar.Value)
}

func TestProgressBar_BindMin(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	done := make(chan bool)
	progressBar := NewProgressBar()
	data := &binding.Float64Binding{}
	progressBar.BindMin(data)
	data.AddListener(func(float64) {
		done <- true
	})
	data.Set(0.75)
	select {
	case <-done:
		time.Sleep(time.Millisecond) // Powernap in case our listener runs first
	case <-time.After(time.Second):
		assert.Fail(t, "Timeout")
	}
	assert.Equal(t, 0.75, progressBar.Min)
}

func TestProgressBar_BindMax(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	done := make(chan bool)
	progressBar := NewProgressBar()
	data := &binding.Float64Binding{}
	progressBar.BindMax(data)
	data.AddListener(func(float64) {
		done <- true
	})
	data.Set(0.75)
	select {
	case <-done:
		time.Sleep(time.Millisecond) // Powernap in case our listener runs first
	case <-time.After(time.Second):
		assert.Fail(t, "Timeout")
	}
	assert.Equal(t, 0.75, progressBar.Max)
}

func TestProgressBar_BindValue(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	done := make(chan bool)
	progressBar := NewProgressBar()
	data := &binding.Float64Binding{}
	progressBar.BindValue(data)
	data.AddListener(func(float64) {
		done <- true
	})
	data.Set(0.75)
	select {
	case <-done:
		time.Sleep(time.Millisecond) // Powernap in case our listener runs first
	case <-time.After(time.Second):
		assert.Fail(t, "Timeout")
	}
	assert.Equal(t, 0.75, progressBar.Value)
}

func TestProgressRenderer_Layout(t *testing.T) {
	bar := NewProgressBar()
	bar.Resize(fyne.NewSize(100, 10))

	render := test.WidgetRenderer(bar).(*progressRenderer)
	assert.Equal(t, 0, render.bar.Size().Width)

	bar.SetValue(.5)
	assert.Equal(t, 50, render.bar.Size().Width)

	bar.SetValue(1)
	assert.Equal(t, bar.Size().Width, render.bar.Size().Width)
}

func TestProgressRenderer_Layout_Overflow(t *testing.T) {
	bar := NewProgressBar()
	bar.Resize(fyne.NewSize(100, 10))

	render := test.WidgetRenderer(bar).(*progressRenderer)
	bar.SetValue(1)
	assert.Equal(t, bar.Size().Width, render.bar.Size().Width)

	bar.SetValue(1.2)
	assert.Equal(t, bar.Size().Width, render.bar.Size().Width)
}

func TestProgressRenderer_ApplyTheme(t *testing.T) {
	bar := NewProgressBar()
	render := test.WidgetRenderer(bar).(*progressRenderer)

	textSize := render.label.TextSize
	customTextSize := textSize
	withTestTheme(func() {
		render.applyTheme()
		customTextSize = render.label.TextSize
	})

	assert.NotEqual(t, textSize, customTextSize)
}
