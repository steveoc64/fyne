package main

// example of a simple list that manages its own content
// but is bound to a List data type.

// uses a simple form binding to simplify the data editor

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/widget"
)

// Customer is a simple set of data against a customer
type Customer struct {
	Name           string  `form:"firstname"`
	Surname        string  `form:"surname"`
	Address        string  `form:"address"`
	Phone          string  `form:"phone"`
	AccountBalance float64 `form:"balance"`
}

// Clone returns a new clone of the existing customer
func (c *Customer) Clone() *Customer {
	newCustomer := Customer(*c)
	return &newCustomer
}

// CustomerList is a simple List type widget that binds
// to a slice of Customer, and manually renders each cell.
//
// We know that there are only a dozen customers in real life
// so we dont need a super optimized list in this case
//
// A manual updating list is good enough for this one
type CustomerListUI struct {
	*widget.Group
	customers *binding.Slice
	binding   *binding.Binding
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
	c.binding = binding.NewBinding(list, c, list)
	list.AddListener(c.binding)
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

// Notify tells the customerList that the dimensions of
// the bound Slice has changed.
// This simple function merely zaps the existing contents
// and rebuilds the world using new widgets.
func (c *CustomerListUI) Notify(b *binding.Binding) {
	c.Clear()
	for i := 0; i < c.customers.Len(); i++ {
		if cust, ok := c.customers.Index(i).(*Customer); ok {
			str := fmt.Sprintf("%s %s, %s, ph %s, bal: $%.2f",
				cust.Name,
				cust.Surname,
				cust.Address,
				cust.Phone,
				cust.AccountBalance,
			)
			c.Append(widget.NewLabel(str))
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
	customerListUI := NewCustomerListUI().Bind(customersData)

	w := a.NewWindow("Customer Editor")

	w.SetContent(widget.NewHBox(
		NewEditorBox(
			widget.NewForm(
				widget.NewFormItem("Name", widget.NewEntry(), "firstname"),
				widget.NewFormItem("Surname", widget.NewEntry(), "surname"),
				widget.NewFormItem("Address", widget.NewEntry(), "address"),
				widget.NewFormItem("Phone", widget.NewEntry(), "phone"),
				widget.NewFormItem("Balance", widget.NewSlider(0, 10000), "balance"),
			).
				Bind(binding.NewStruct(editCustomer)).
				OnSubmit(func() {
					// we need to clone the customer and append it to the list, then clear the form
					customersData.Append(editCustomer.Clone())
					editBinding.SetValue(&Customer{})
				}),
			widget.NewButton("Show List", func() {
				// example of being able to do useful things with the
				// binding.List outside of widgets and outside of
				// the data model
				println("Customer List")
				println("==============================")
				for i := 0; i < customersData.Len(); i++ {
					if c, ok := customersData.Index(i).(*Customer); ok {
						println(i+1, ":", c.Name, c.Surname, c.Address, c.Phone, c.AccountBalance)
					}
				}
			}),
		),
		customerListUI,
	))
	w.ShowAndRun()
}
