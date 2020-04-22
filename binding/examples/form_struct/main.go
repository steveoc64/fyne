// example of binding a form to a struct
package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/widget"
)

type Customer struct {
	Title          string  `form:"title"`
	Name           string  `form:"name"`
	Address        string  `form:"address"`
	Phone          string  `form:"phone"`
	VIP            bool    `form:"vip"`
	AccountBalance float64 `form:"balance"`
}

func main() {
	a := app.New()

	customer := &Customer{
		Title:   "Mr",
		Name:    "Fred",
		Address: "1 main street",
		Phone:   "1234",
	}

	custBinding := binding.NewStruct(customer)

	w := a.NewWindow("StructBinding")
	form1 := widget.NewForm(
		widget.NewFormItem("Name", widget.NewEntry(), "name"),
		widget.NewFormItem("Title", widget.NewEntry(), "title"),
		widget.NewFormItem("", widget.NewCheck("VIP Customer", nil), "vip"),
		widget.NewFormItem("Address", widget.NewEntry(), "address"),
		widget.NewFormItem("Phone", widget.NewEntry(), "phone"),
		widget.NewFormItem("Balance", widget.NewSlider(0, 10000), "balance"),
	).Bind(custBinding)

	w.SetContent(
		widget.NewGroup("Customer Details View 1",
			form1,
			widget.NewButton("Submit", func() {
				form1.Submit()
			}),
		),
	)
	w.Show()

	w2 := a.NewWindow("StructBinding")
	form2 := widget.NewForm(
		widget.NewFormItem("", widget.NewCheck("VIP Customer", nil), "vip"),
		widget.NewFormItem("Title", widget.NewEntry(), "title"),
		widget.NewFormItem("Name", widget.NewEntry(), "name"),
		widget.NewFormItem("Balance", widget.NewSlider(0, 10000), "balance"),
		widget.NewFormItem("Phone", widget.NewEntry(), "phone"),
		widget.NewFormItem("Address", widget.NewEntry(), "address"),
	).Bind(custBinding)

	w2.SetContent(
		widget.NewGroup("Customer View 2",
			form2,
			widget.NewButton("Submit", func() {
				form2.Submit()
			}),
		),
	)
	w2.ShowAndRun()
}
