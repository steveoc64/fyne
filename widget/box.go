package widget

import (
	"image/color"
	"sync"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
)

// Box widget is a simple list where the child elements are arranged in a single column
// for vertical or a single row for horizontal arrangement
type Box struct {
	sync.RWMutex
	baseWidget
	background color.Color

	Horizontal bool
	Children   []fyne.CanvasObject
}

// Resize sets a new size for a widget.
// Note this should not be used if the widget is being managed by a Layout within a Container.
func (b *Box) Resize(size fyne.Size) {
	b.resize(size, b)
}

// Move the widget to a new position, relative to it's parent.
// Note this should not be used if the widget is being managed by a Layout within a Container.
func (b *Box) Move(pos fyne.Position) {
	b.move(pos, b)
}

// MinSize returns the smallest size this widget can shrink to
func (b *Box) MinSize() fyne.Size {
	return b.minSize(b)
}

// Show this widget, if it was previously hidden
func (b *Box) Show() {
	b.show(b)
}

// Hide this widget, if it was previously visible
func (b *Box) Hide() {
	b.hide(b)
}

// ApplyTheme updates this box to match the current theme
func (b *Box) ApplyTheme() {
	b.Lock()
	b.background = theme.BackgroundColor()
	b.Unlock()
}

// Prepend inserts a new CanvasObject at the top/left of the box
func (b *Box) Prepend(object fyne.CanvasObject) {
	b.Lock()
	b.Children = append([]fyne.CanvasObject{object}, b.Children...)
	b.Unlock()

	Refresh(b)
}

// Append adds a new CanvasObject to the end/right of the box
func (b *Box) Append(object fyne.CanvasObject) {
	b.Lock()
	b.Children = append(b.Children, object)
	b.Unlock()

	Refresh(b)
}

// CreateRenderer is a private method to Fyne which links this widget to it's renderer
func (b *Box) CreateRenderer() fyne.WidgetRenderer {
	b.RLock()
	defer b.RUnlock()
	var lay fyne.Layout
	if b.Horizontal {
		lay = layout.NewHBoxLayout()
	} else {
		lay = layout.NewVBoxLayout()
	}

	return &boxRenderer{objects: b.Children, layout: lay, box: b}
}

func (b *Box) setBackgroundColor(bg color.Color) {
	b.Lock()
	b.background = bg
	b.Unlock()
}

// NewHBox creates a new horizontally aligned box widget with the specified list of child objects
func NewHBox(children ...fyne.CanvasObject) *Box {
	box := &Box{baseWidget: baseWidget{}, Horizontal: true, Children: children}

	Renderer(box).Layout(box.MinSize())
	return box
}

// NewVBox creates a new vertically aligned box widget with the specified list of child objects
func NewVBox(children ...fyne.CanvasObject) *Box {
	box := &Box{baseWidget: baseWidget{}, Horizontal: false, Children: children}

	Renderer(box).Layout(box.MinSize())
	return box
}

type boxRenderer struct {
	sync.RWMutex
	layout fyne.Layout

	objects []fyne.CanvasObject
	box     *Box
}

func (b *boxRenderer) MinSize() fyne.Size {
	b.RLock()
	defer b.RUnlock()
	return b.layout.MinSize(b.objects)
}

func (b *boxRenderer) Layout(size fyne.Size) {
	b.RLock()
	b.layout.Layout(b.objects, size)
	b.RUnlock()
}

func (b *boxRenderer) ApplyTheme() {
}

func (b *boxRenderer) BackgroundColor() color.Color {
	b.box.RLock()
	defer b.box.RUnlock()
	if b.box.background == nil {
		return theme.BackgroundColor()
	}

	return b.box.background
}

func (b *boxRenderer) Objects() []fyne.CanvasObject {
	return b.objects
}

func (b *boxRenderer) Refresh() {
	b.Lock()
	b.box.RLock()
	b.objects = b.box.Children
	b.Unlock()
	b.Layout(b.box.Size())
	b.box.RUnlock()

	canvas.Refresh(b.box)
}

func (b *boxRenderer) Destroy() {
}
