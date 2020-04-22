package widget

import (
	"errors"
	"fmt"
	"reflect"

	"fyne.io/fyne"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/internal/cache"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"github.com/davecgh/go-spew/spew"
)

// FormItem provides the details for a row in a form
type FormItem struct {
	tag    string
	Text   string
	Widget fyne.CanvasObject
}

// NewFormItem creates a new form item with the specified label text and input widget
func NewFormItem(text string, widget fyne.CanvasObject, tag string) *FormItem {
	return &FormItem{tag, text, widget}
}

// Form widget is two column grid where each row has a label and a widget (usually an input).
// The last row of the grid will contain the appropriate form control buttons if any should be shown.
// Setting OnSubmit will set the submit button to be visible and call back the function when tapped.
// Setting OnCancel will do the same for a cancel button.
type Form struct {
	BaseWidget

	Items      []*FormItem
	OnSubmit   func()
	OnCancel   func()
	SubmitText string
	CancelText string

	onSubmit func()
	onCancel func()

	itemGrid *fyne.Container
	binding  *binding.Binding
}

func (f *Form) createLabel(text string) *Label {
	return NewLabelWithStyle(text, fyne.TextAlignTrailing, fyne.TextStyle{Bold: true})
}

func (f *Form) ensureGrid() {
	if f.itemGrid != nil {
		return
	}

	f.itemGrid = fyne.NewContainerWithLayout(layout.NewFormLayout(), []fyne.CanvasObject{}...)
}

// Append adds a new row to the form, using the text as a label next to the specified Widget
func (f *Form) Append(text string, widget fyne.CanvasObject) *Form {
	item := &FormItem{Text: text, Widget: widget}
	f.AppendItem(item)
	return f
}

// AppendItem adds the specified row to the end of the Form
func (f *Form) AppendItem(items ...*FormItem) *Form {
	f.ExtendBaseWidget(f) // could be called before render

	// ensure we have a renderer set up (that creates itemGrid)...
	cache.Renderer(f.super())

	for _, item := range items {
		f.Items = append(f.Items, item)
		f.itemGrid.AddObject(f.createLabel(item.Text))
		f.itemGrid.AddObject(item.Widget)
	}

	f.Refresh()
	return f
}

// MinSize returns the size that this widget should not shrink below
func (f *Form) MinSize() fyne.Size {
	f.ExtendBaseWidget(f)
	return f.BaseWidget.MinSize()
}

// Refresh updates the widget state when requested.
func (f *Form) Refresh() {
	f.BaseWidget.Refresh()
	canvas.Refresh(f) // refresh ourselves for BG color - the above updates the content
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (f *Form) CreateRenderer() fyne.WidgetRenderer {
	f.ExtendBaseWidget(f)
	f.ensureGrid()
	for _, item := range f.Items {
		f.itemGrid.AddObject(f.createLabel(item.Text))
		f.itemGrid.AddObject(item.Widget)
	}

	if f.OnCancel == nil && f.OnSubmit == nil {
		return cache.Renderer(NewVBox(f.itemGrid))
	}

	buttons := NewHBox(layout.NewSpacer())
	if f.OnCancel != nil {
		if f.CancelText == "" {
			f.CancelText = "Cancel"
		}

		buttons.Append(NewButtonWithIcon(f.CancelText, theme.CancelIcon(), f.OnCancel))
	}
	if f.OnSubmit != nil {
		if f.SubmitText == "" {
			f.SubmitText = "Submit"
		}

		submitButton := NewButtonWithIcon(f.SubmitText, theme.ConfirmIcon(), f.OnSubmit)
		submitButton.Style = PrimaryButton
		buttons.Append(submitButton)
	}
	return cache.Renderer(NewVBox(f.itemGrid, buttons))
}

// NewForm creates a new form widget with the specified rows of form items
// and (if any of them should be shown) a form controls row at the bottom
func NewForm(items ...*FormItem) *Form {
	form := &Form{BaseWidget: BaseWidget{}, Items: items}
	form.ExtendBaseWidget(form)

	return form
}

// OnSubmit sets an on submit handler for the form
func (f *Form) OnSubmit(f func()) *Form {
	f.onSubmit = f
	return f
}

// OnCancel sets an on cancel handler for the form
func (f *Form) OnCancel(f func()) *Form {
	f.onCancel = f
	return f
}

// Bind creates a new Binding between this form and some struct
func (f *Form) Bind(value binding.Data) *Form {
	if value == nil {
		// invalid - return
		return f
	}
	b := binding.NewBinding(value, f, value)
	if b.Handler.Kind() != reflect.Struct {
		// invalid binding type, do not connect
		return f
	}
	f.binding = b
	return f
}

// Handler (optionally) sets the handler for this binding
func (f *Form) Handler(h binding.Handler) *Form {
	if f.binding != nil {
		if h.Kind() != reflect.Struct {
			// invalid, do not chain with this
			fyne.LogError("Form.Handler()", errors.New("form handlers must be of type binding.Struct"))
			return f
		}
		f.binding.Handler = h
	}
	return f
}

// Unbind disconnects the widget and puts the binding out for garbage collection
func (f *Form) Unbind() *Form {
	if f.binding != nil {
		f.binding.Data.DeleteListener(f.binding)
		f.binding = nil
	}
	return f
}

// Notify tells the form that the struct that its bound to has been updated
// so pull the data and populate the form fields
func (f *Form) Notify(b *binding.Binding) {
	if f == nil {
		// is actually possible, so trap it here
		return
	}
	// we can get the data from the binding
	// we know that the handler always returns a string value
	value := b.Handler.Get()
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		fyne.LogError("Form.Notify()", fmt.Errorf("Get() returned %T, expecting struct", value.Type().Name()))
	}

	// iterate through the struct, and send the data to each field in the form
	dataType := value.Type()
	for i := 0; i < value.NumField(); i++ {
		fld := dataType.Field(i)
		tag := fld.Tag.Get("form")
		v := value.Field(i)
		fmt.Printf("%d: %s = %v\n", i, tag, v)
		if element, ok := f.Element(tag); ok {
			switch element.(type) {
			case *Entry:
				element.(*Entry).SetText(v.String())
			case *Check:
				element.(*Check).SetChecked(v.Bool())
			case *Slider:
				element.(*Slider).SetValue(v.Float())
			}
		}
	}
}

// Element gets the element with the matching tag name
func (f *Form) Element(tag string) (fyne.CanvasObject, bool) {
	for _, item := range f.Items {
		if item.tag == tag {
			return item.Widget, true
		}
	}
	return nil, false
}

// Submit will fill in the bound struct from the contents of the fields
func (f *Form) Submit() {
	// TODO :
	// - validation hook
	// - pre commit hook
	// - post commit hook

	if f.binding == nil {
		return
	}
	value := f.binding.Handler.Get()
	savePtr := value
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	dataType := value.Type()
	for i := 0; i < value.NumField(); i++ {
		fld := dataType.Field(i)
		tag := fld.Tag.Get("form")
		v := value.Field(i)
		fmt.Printf("%d: %s = %v\n", i, tag, v)
		if element, ok := f.Element(tag); ok {
			switch element.(type) {
			case *Entry:
				v.SetString(element.(*Entry).Text)
			case *Check:
				v.SetBool(element.(*Check).Checked)
			case *Slider:
				v.SetFloat(element.(*Slider).Value)
			}
		}
		spew.Dump("set field", i, v)
	}
	spew.Dump("set data", value)
	spew.Dump("from ptr", savePtr)
	f.binding.Data.Update()
}
