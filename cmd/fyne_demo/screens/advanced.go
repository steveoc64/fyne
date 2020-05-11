package screens

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/theme"
)

func scaleString(c fyne.Canvas) string {
	return fmt.Sprintf("%0.2f", c.Scale())
}

func prependTo(g *widget.Group, s string) {
	g.Prepend(widget.NewLabel(s))
}

func setScaleText(obj *widget.Label, win fyne.Window) {
	for obj.Visible() {
		obj.SetText(scaleString(win.Canvas()))

		time.Sleep(time.Second)
	}
}

// AdvancedScreen loads a panel that shows details and settings that are a bit
// more detailed than normally needed.
func AdvancedScreen(win fyne.Window) fyne.CanvasObject {
	scale := widget.NewLabel("")

	screen := widget.NewGroup("Screen", widget.NewForm(
		&widget.FormItem{Text: "Scale", Widget: scale},
	))

	go setScaleText(scale, win)

	label := widget.NewLabel("Just type...")
	generic := widget.NewGroupWithScroller("Generic")
	desk := widget.NewGroupWithScroller("Desktop")

	win.Canvas().SetOnTypedRune(func(r rune) {
		prependTo(generic, "Rune: "+string(r))
	})
	win.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		prependTo(generic, "Key : "+string(ev.Name))
	})
	if deskCanvas, ok := win.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(ev *fyne.KeyEvent) {
			prependTo(desk, "KeyDown: "+string(ev.Name))
		})
		deskCanvas.SetOnKeyUp(func(ev *fyne.KeyEvent) {
			prependTo(desk, "KeyUp  : "+string(ev.Name))
		})
	}

	themeUpdateEntry := widget.NewMultiLineEntry()
	return widget.NewHBox(
		widget.NewVBox(screen,
			themeUpdateEntry,
			widget.NewButton("Apply Custom Theme", func() {
				if err := theme.Extend(themeUpdateEntry.Text); err != nil {
					println("Error:", err.Error())
				}
				themeUpdateEntry.SetText("")

			}),
			widget.NewButton("Fullscreen", func() {
				win.SetFullScreen(!win.FullScreen())
			}),
		),
		fyne.NewContainerWithLayout(layout.NewBorderLayout(label, nil, nil, nil),
			label,
			fyne.NewContainerWithLayout(layout.NewGridLayout(2),
				generic, desk,
			),
		),
	)
}
