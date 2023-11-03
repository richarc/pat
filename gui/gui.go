package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// PatApp is a structure representing the core elements of the Pat application
type PatApp struct {
	App        fyne.App
	MainWindow fyne.Window
}

var Pat PatApp

// Initialise the key elements of the GUI
// The App and window are defined here, the run loop is executed in the main function
func init() {
	Pat.App = app.New()
	Pat.MainWindow = Pat.App.NewWindow("Main Window")
	Pat.MainWindow.SetContent(Pat.MakeUI())
	Pat.MainWindow.Resize(fyne.NewSize(800, 800))
}

func (p PatApp) MakeUI() fyne.CanvasObject {

	// list of options for the selct GET, POST etc..
	actions := []string{"GET", "POST"}
	sel := widget.NewSelect(actions, func(value string) {
		log.Println("Select set to", value)
	})

	//create a radio group to select HTTP or HTTPS
	radio := widget.NewRadioGroup([]string{"HTTP", "HTTPS"}, func(value string) {
		log.Println("Radio set to", value)
	})

	//The Host entry field
	host_entry := widget.NewEntry()
	host_entry.SetPlaceHolder("Enter the host name or IP Address...")

	//The Host entry field
	uri_entry := widget.NewEntry()
	uri_entry.SetPlaceHolder("Enter the URI and query string...")

	main_container := container.NewVBox(sel, radio, host_entry, uri_entry)
	return main_container
}
