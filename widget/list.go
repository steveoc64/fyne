package widget

import (
	"image/color"
	"log"
	"reflect"
	"sync"

	"fyne.io/fyne"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/internal/cache"
	"fyne.io/fyne/theme"
)

var (
	listMinSize    = 32 // TODO consider the smallest useful list view?
	listMinVisible = 1.75
)

// List widget is a view of a slice of some constant height canvasObjects
// assumes vertical orientation for the default list widget
type List struct {
	BaseWidget
	cellSize fyne.Size
	disabled bool
	hidden   bool
	padding  int
	pos      fyne.Position
	selected int
	size     fyne.Size

	index        int // Index of first visible item
	offsetItem   int // Offset from top-left of visible area to top-left of cell.
	offsetScroll int // Offset from top-left of visible area to top-left of widget.

	// TODO Add ScrollBars
	minSize  fyne.Size
	minSizes []fyne.Size

	// binding options
	isBound        bool
	binding        binding.Slice
	listHandler    binding.Handler
	elementType    reflect.Type
	elementHandler binding.Handler
}

// NewList creates a new list widget with the set items and change handler
func NewList(items binding.Slice) *List {
	return &List{}
}

// Bind binds the list to a data source of type slice
func (l *List) Bind(items binding.Slice) *List {
	l.binding = items
	l.elementType = items.ElementType()
	l.isBound = true
	return l
}

// Handler allows the user to override the handler to use in accessing the bound data
// by default the item's default handler is used
func (l *List) Handler(h binding.Handler) *List {
	l.listHandler = h
	return l
}

// ElementHandler allows the user to override the handler to be used
// when accessing the elements in the list
func (l *List) ELementHandler(h binding.Handler) *List {
	l.elementHandler = h
	return l
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (l *List) CreateRenderer() fyne.WidgetRenderer {
	log.Println("List.CreateRenderer")
	r := &listRenderer{
		list: l,
	}
	return r
}

// Enable this widget, updating any style or features appropriately.
func (l *List) Enable() {
	l.disabled = false
}

// Disable this widget so that it cannot be interacted with, updating any style appropriately.
func (l *List) Disable() {
	l.disabled = true
}

// Hide this widget so it is no lonver visible
func (l *List) Hide() {
	l.hidden = true
}

// IsDisabled returns true if this widget is currently disabled or false if it can currently be interacted with.
func (l *List) IsDisabled() (disabled bool) {
	return l.disabled
}

// MinSize returns the size that this widget should not shrink below
func (l *List) MinSize() fyne.Size {
	min := cache.Renderer(l).(*listRenderer).MinSize()
	min.Height = int(float64(min.Height) * listMinVisible)
	l.minSize = l.minSize.Max(min)
	log.Println("List.MinSize:", l.minSize)
	return l.minSize
}

// MouseIn is called when a desktop pointer enters the widget
func (l *List) MouseIn(event *desktop.MouseEvent) {
	if l.IsDisabled() {
		return
	}
	// log.Println("List.MouseIn:", event)
	// TODO
}

func (l *List) Show() {
	l.hidden = false
}

// MouseOut is called when a desktop pointer exits the widget
func (l *List) MouseOut() {
	if l.IsDisabled() {
		return
	}
	// log.Println("List.MouseOut")
	// TODO
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (l *List) MouseMoved(event *desktop.MouseEvent) {
	if l.IsDisabled() {
		return
	}
	// log.Println("List.MouseMoved:", event)
	// TODO
}

// Move the widget to a new position, relative to its parent.
func (l *List) Move(pos fyne.Position) {
	log.Println("List.Move:", pos)
	l.pos = pos
}

// Position gets the current position of this widget, relative to its parent.
func (l *List) Position() fyne.Position {
	return l.pos
}

// Refresh causes this widget to be redrawn in it's current state
func (l *List) Refresh() {
	log.Println("List.Refresh")
	// no-op
}

// Resize sets a new size for a widget.
func (l *List) Resize(size fyne.Size) {
	log.Println("List.Resize:", size)
	l.size = size
}

// Scrolled is called when an input device triggers a scroll event
func (l *List) Scrolled(event *fyne.ScrollEvent) {
	log.Println("List.Scrolled:", event)
}

// Size gets the current size of this widget.
func (l *List) Size() (size fyne.Size) {
	return l.size
}

// Tapped is called when a pointer tapped event is captured and triggers any change handler
func (l *List) Tapped(event *fyne.PointEvent) {
	if l.IsDisabled() {
		return
	}
}

// Visible returns whether or not this widget should be visible.
func (l *List) Visible() bool {
	return !l.hidden
}

type listElement struct {
	fyne.CanvasObject
	bound binding.Notifiable
}

type listRenderer struct {
	list  *List
	cells []listElement
	pool  sync.Pool
}

// BackgroundColor satisfies the fyne.WidgetRenderer interface.
func (r *listRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

// Destroy satisfies the fyne.WidgetRenderer interface.
func (r *listRenderer) Destroy() {
}

// Layout the visible items of the list widget
func (r *listRenderer) Layout(size fyne.Size) {
}

// MinSize calculates the largest minimum size of the visible list items.
func (r *listRenderer) MinSize() (size fyne.Size) {
	size.Width = listMinSize
	size.Height = listMinSize
	for _, c := range r.cells {
		size = size.Max(c.MinSize())
	}
	// log.Println("listRenderer.MinSize:", size)
	return
}

// Objects satisfies the fyne.WidgetRenderer interface.
func (r *listRenderer) Objects() []fyne.CanvasObject {
	obj := make([]fyne.CanvasObject, len(r.cells))
	for i, v := range r.cells {
		obj[i] = v.CanvasObject
	}
	return obj
}

// Refresh satisfies the fyne.WidgetRenderer interface.
func (r *listRenderer) Refresh() {
	// log.Println("listRenderer.Refresh")
	// no-op
}

func (r *listRenderer) updateItems(start, length int) {
	indexCell := 0
	indexItem := start
	x := -r.list.offsetItem
	y := -r.list.offsetItem
	var cellSize fyne.Size
	var c listElement
	var min fyne.Size
	for ; indexItem < length && x < r.list.size.Width && y < r.list.size.Height; indexCell, indexItem = indexCell+1, indexItem+1 {
		if indexCell < len(r.cells) {
			c = r.cells[indexCell]
		} else {
			c = r.pool.Get().(listElement)
			r.cells = append(r.cells, c)
		}

		// if using binding, then trigger a notify on the element
		if r.list.isBound {
			data := r.list.binding.Index(indexItem) // get whatever the slice we are bound to is storing at that index
			c.bound.Notify(&binding.Binding{
				Data:    data.(binding.Observable),
				Element: c.bound,
				Handler: r.list.elementHandler,
			})
		}

		if cellSize.IsZero() {
			min = c.MinSize()
			// log.Println(indexItem, min)
			if indexItem < len(r.list.minSizes) {
				r.list.minSizes[indexItem] = min
			} else {
				r.list.minSizes = append(r.list.minSizes, min)
			}
		} else {
			min = cellSize
		}
		y += min.Height
	}
	// Recycle unused cells
	lastCell := indexCell
	for ; indexCell < len(r.cells); indexCell++ {
		c := r.cells[indexCell]
		c.Hide()
		r.pool.Put(c)
	}
	r.cells = r.cells[:lastCell]
}
