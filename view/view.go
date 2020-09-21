package view

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var application fyne.App
var window fyne.Window

func test() {
	fmt.Println("Tapped")
}


func Init() {
	application =  app.New() 
	window = application.NewWindow("Hello")
	hello := widget.NewLabel("Hello Fyne!")
	window.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
		}),
	))


	//newButton := widget.NewButton("hello", test)

	//window.SetContent(newButton)
	
	window.Resize(fyne.NewSize(600,600))
	window.ShowAndRun()
}

