package main

import (
	"fmt"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dataapi"
	_ "image/png"

	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"fyne.io/fyne"
)

// View is a window with reactive elements
type View struct {
	app  fyne.App
	data *dataModel
	w    fyne.Window
}

func newView(app fyne.App, data *dataModel, logo *canvas.Image) *View {
	data.NumWindows++
	title := fmt.Sprintf("View #%d", data.NumWindows)

	v := &View{
		app:  app,
		data: data,
		w:    app.NewWindow(title),
	}

	myWindowID := data.NumWindows
	v.w.SetContent(widget.NewVBox(
		widget.NewLabelWithStyle("Reactive Data", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		logo,
		layout.NewSpacer(),
		widget.NewForm(
			widget.NewFormItem("Time", dataapi.NewLabel(data.Time)),
			widget.NewFormItem("Name", dataapi.NewEntry(data.Name)),
			widget.NewFormItem("", dataapi.NewCheck(data.IsAvailable,
				"Avail",
				func(checked bool) {
					println("clicked on the avail button", checked)
				})),
			widget.NewFormItem("", widget.NewRadio([]string{"Small", "Medium", "Large"}, func(value string) {
				println("Radio button changed to", value)
			})),
			widget.NewFormItem("Delivery", widget.NewSlider(0.0, 100.0)),
			widget.NewFormItem("Quit Cleanly", widget.NewButton("Quit", func() {
				data.Time.DeleteListener(myWindowID)
				v.w.Close()
			})),
		),
	))

	v.w.Show()
	return v
}
