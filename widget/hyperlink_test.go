package widget

import (
	"net/url"
	"testing"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/test"
	"fyne.io/fyne/theme"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHyperlink_MinSize(t *testing.T) {
	u, err := url.Parse("https://fyne.io/")
	assert.Nil(t, err)

	hyperlink := NewHyperlink("Test", u)
	min := hyperlink.MinSize()

	assert.True(t, min.Width > theme.Padding()*2)

	hyperlink.SetText("Longer")
	assert.True(t, hyperlink.MinSize().Width > min.Width)
}

func TestHyperlink_Cursor(t *testing.T) {
	u, err := url.Parse("https://fyne.io/")
	hyperlink := NewHyperlink("Test", u)

	assert.Nil(t, err)
	assert.Equal(t, desktop.PointerCursor, hyperlink.Cursor())
}

func TestHyperlink_Alignment(t *testing.T) {
	hyperlink := &Hyperlink{Text: "Test", Alignment: fyne.TextAlignTrailing}
	assert.Equal(t, fyne.TextAlignTrailing, textRenderTexts(hyperlink)[0].Alignment)
}

func TestHyperlink_SetText(t *testing.T) {
	u, err := url.Parse("https://fyne.io/")
	assert.Nil(t, err)

	hyperlink := &Hyperlink{Text: "Test", URL: u}
	hyperlink.Refresh()
	hyperlink.SetText("New")

	assert.Equal(t, "New", hyperlink.Text)
	assert.Equal(t, "New", textRenderTexts(hyperlink)[0].Text)
}

func TestHyperlink_SetUrl(t *testing.T) {
	sURL, err := url.Parse("https://github.com/fyne-io/fyne")
	assert.Nil(t, err)

	// test constructor
	hyperlink := NewHyperlink("Test", sURL)
	assert.Equal(t, sURL, hyperlink.URL)

	// test setting functions
	sURL, err = url.Parse("https://fyne.io")
	assert.Nil(t, err)
	hyperlink.SetURL(sURL)
	assert.Equal(t, sURL, hyperlink.URL)
}

func TestHyperlink_BindText(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	done := make(chan bool)
	u, err := url.Parse("https://fyne.io")
	assert.Nil(t, err)
	hyperlink := NewHyperlink("hyperlink", u)
	data := &binding.StringBinding{}
	hyperlink.BindText(data)
	data.AddListenerFunction(func(binding.Binding) {
		done <- true
	})
	data.Set("foobar")
	select {
	case <-done:
		time.Sleep(time.Millisecond) // Powernap in case our listener runs first
	case <-time.After(time.Second):
		assert.Fail(t, "Timeout")
	}
	assert.Equal(t, "foobar", hyperlink.Text)
}

func TestHyperlink_BindURL(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	done := make(chan bool)
	u, err := url.Parse("https://fyne.io")
	assert.Nil(t, err)
	hyperlink := NewHyperlink("hyperlink", u)
	data := &binding.URLBinding{}
	hyperlink.BindURL(data)
	u, err = url.Parse("https://github.com/fyne-io/fyne")
	assert.Nil(t, err)
	data.AddListenerFunction(func(binding.Binding) {
		done <- true
	})
	data.Set(u)
	select {
	case <-done:
		time.Sleep(time.Millisecond) // Powernap in case our listener runs first
	case <-time.After(time.Second):
		assert.Fail(t, "Timeout")
	}
	assert.Equal(t, u, hyperlink.URL)
}

func TestHyperlink_CreateRendererDoesNotAffectSize(t *testing.T) {
	u, err := url.Parse("https://github.com/fyne-io/fyne")
	require.NoError(t, err)
	link := NewHyperlink("Test", u)
	link.Resize(link.MinSize())
	size := link.Size()
	assert.NotEqual(t, fyne.NewSize(0, 0), size)
	assert.Equal(t, size, link.MinSize())

	r := link.CreateRenderer()
	assert.Equal(t, size, link.Size())
	assert.Equal(t, size, link.MinSize())
	assert.Equal(t, size, r.MinSize())
	r.Layout(size)
	assert.Equal(t, size, link.Size())
	assert.Equal(t, size, link.MinSize())
}
