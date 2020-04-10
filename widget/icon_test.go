package widget

import (
	"testing"
	"time"

	"fyne.io/fyne/binding"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/test"
	"fyne.io/fyne/theme"
	"github.com/stretchr/testify/assert"
)

func TestNewIcon(t *testing.T) {
	icon := NewIcon(theme.ConfirmIcon())
	render := test.WidgetRenderer(icon)

	assert.Equal(t, 1, len(render.Objects()))
	obj := render.Objects()[0]
	img, ok := obj.(*canvas.Image)
	if !ok {
		t.Fail()
	}
	assert.Equal(t, theme.ConfirmIcon(), img.Resource)
}

func TestIcon_Nil(t *testing.T) {
	icon := NewIcon(nil)
	render := test.WidgetRenderer(icon)

	assert.Equal(t, 0, len(render.Objects()))
}

func TestIcon_MinSize(t *testing.T) {
	icon := NewIcon(theme.CancelIcon())
	min := icon.MinSize()

	assert.Equal(t, theme.IconInlineSize(), min.Width)
	assert.Equal(t, theme.IconInlineSize(), min.Height)
}

func TestIcon_BindResource(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	done := make(chan bool)
	icon := NewIcon(theme.WarningIcon())
	data := &binding.ResourceBinding{}
	icon.BindResource(data)
	data.AddListenerFunction(func(binding.Binding) {
		done <- true
	})
	data.Set(theme.InfoIcon())
	select {
	case <-done:
		time.Sleep(time.Millisecond) // Powernap in case our listener runs first
	case <-time.After(time.Second):
		assert.Fail(t, "Timeout")
	}
	assert.Equal(t, theme.InfoIcon(), icon.Resource)
}

func TestIconRenderer_ApplyTheme(t *testing.T) {
	icon := NewIcon(theme.CancelIcon())
	render := test.WidgetRenderer(icon).(*iconRenderer)
	visible := render.objects[0].Visible()

	render.Refresh()
	assert.Equal(t, visible, render.objects[0].Visible())
}
