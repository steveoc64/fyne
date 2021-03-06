package main

//go:generate fyne bundle reactive-data.png > image.go

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

const preferenceCurrentTab = "currentTab"

func welcomeScreen(a fyne.App, data *dataModel, diag, logo *canvas.Image) fyne.CanvasObject {
	return widget.NewVBox(
		widget.NewLabelWithStyle("Fyne Reactive Data Update Demo", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewHBox(
			layout.NewSpacer(),
			diag,
			layout.NewSpacer(),
		),
		widget.NewLabel(`This demo has a single instance of a dataapi model as described above.

Each window subscribes to the dataapi model, and is automatically updated
when the model changes.

Changes to the dataapi in the view are committed to the central dataapi model,
which in turn triggers a repaint on the other subscribed views.`),
		layout.NewSpacer(),
		widget.NewGroup("Data Model Internals",
			fyne.NewContainerWithLayout(layout.NewGridLayout(4),
				widget.NewLabel("Clock"),
				widget.NewLabelWithData(data.Clock),
				widget.NewLabel("Name (String)"),
				widget.NewLabelWithData(data.Name),
				widget.NewLabel("Avail (Bool)"),
				widget.NewLabelWithData(data.IsAvailable),
				widget.NewLabel("OnSale (String)"),
				widget.NewLabelWithData(data.OnSale),
				widget.NewLabel("Size (Int)"),
				widget.NewLabelWithData(data.Size),
				widget.NewLabel("Deliv Time (Float)"),
				widget.NewLabelWithData(data.DeliveryTime),
			),
		),
		widget.NewGroup("Launch New Viewer",
			fyne.NewContainerWithLayout(layout.NewGridLayout(5),
				layout.NewSpacer(),
				layout.NewSpacer(),
				widget.NewButton("+ Viewer Window", func() {
					newView(a, data, logo)
				}),
				widget.NewButton("Fx1", func() {
					newFx(a, data)
				}),
				widget.NewButton("Fx2", func() {
					newPic(a, data)
				}),
				layout.NewSpacer(),
				widget.NewButton("Close", func() { a.Quit() }),
			),
		),
	)
}

func main() {
	logo := canvas.NewImageFromResource(theme.FyneLogo())
	logo.SetMinSize(fyne.NewSize(320, 128))
	diag := canvas.NewImageFromResource(resourceReactiveDataPng)
	diag.SetMinSize(fyne.NewSize(400, 350))

	data := NewDataModel()
	a := app.NewWithID("io.fyne.reactive")
	w := a.NewWindow("Fyne Reactive Demo")
	w.SetMaster()

	w.SetContent(welcomeScreen(a, data, diag, logo))

	w.ShowAndRun()
}
