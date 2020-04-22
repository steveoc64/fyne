// example of a simple list that manages its own content
// but is bound to a List data type.
package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/widget"
)

// Customer is a simple set of data against a customer
type Customer struct {
	Name    string  `form:"firstname"`
	Surname string  `form:"surname"`
	Phone   string  `form:"phone"`
	Address string  `form:"address"`
	Balance float64 `form:"balance"`
}

// Clone returns a new clone of the existing customer
func (c *Customer) Clone() *Customer {
	newCustomer := Customer(*c)
	return &newCustomer
}

func (c *Customer) String() string {
	return fmt.Sprintf("%s %s %s ph: %s, Bal $%.02f",
		c.Name,
		c.Surname,
		c.Address,
		c.Phone,
		c.Balance,
	)
}

// CustomerList is a simple List type widget that binds
// to a slice of Customer, and manually renders each cell.
type CustomerListUI struct {
	*widget.Group
	customers  *binding.Slice
	connection *binding.Binding
}

// NewCustomerList retuns a new CustomerList UI Element
func NewCustomerListUI() *CustomerListUI {
	c := &CustomerListUI{
		Group: widget.NewGroupWithScroller("Customer List"),
	}
	return c
}

func (c *CustomerListUI) Bind(list *binding.Slice) *CustomerListUI {
	c.customers = list
	c.connection = binding.Connect(list, c, list)
	list.AddListener(c.connection)
	return c
}

// MinSize hardcoded for now
func (c *CustomerListUI) MinSize() fyne.Size {
	return fyne.NewSize(600, 20)
}

type EditorBox struct {
	*widget.Box
}

func NewEditorBox(children ...fyne.CanvasObject) *EditorBox {
	return &EditorBox{
		widget.NewVBox(children...),
	}
}

func (c *EditorBox) MinSize() fyne.Size {
	return fyne.NewSize(300, 300)
}

// Notify tells the customerListUI that the dimensions of
// the bound Slice has changed.
func (c *CustomerListUI) Notify(b *binding.Binding) {
	c.Clear()
	for i := 0; i < c.customers.Len(); i++ {
		if cust, ok := c.customers.Index(i).(*Customer); ok {
			c.Append(widget.NewLabel(cust.String()))
		}
	}
}

func main() {
	// setup our data model with binding
	editCustomer := &Customer{}
	editBinding := binding.NewStruct(editCustomer)
	customersData := binding.NewSlice(&Customer{}, 0, 1000)

	// create the app
	a := app.New()

	// build the UI

	w := a.NewWindow("Customer Editor")
	w.SetContent(widget.NewHBox(
		NewEditorBox(
			widget.NewForm(
				widget.NewFormItem("Name", widget.NewEntry(), "firstname"),
				widget.NewFormItem("Surname", widget.NewEntry(), "surname"),
				widget.NewFormItem("Address", widget.NewEntry(), "address"),
				widget.NewFormItem("Phone", widget.NewEntry(), "phone"),
				widget.NewFormItem("Balance", widget.NewSlider(0, 1000), "balance"),
			).
				Bind(binding.NewStruct(editCustomer)).
				OnSubmit(func() {
					// we need to clone the customer and append it to the list, then clear the form
					customersData.Append(editCustomer.Clone())
					editBinding.SetValue(&Customer{})
				}),
		),
		NewCustomerListUI().Bind(customersData)),
	)
	w.ShowAndRun()
}
