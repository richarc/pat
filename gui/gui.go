package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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

// MakeUI create the canvas object for the main window
// it calls two additional functions to create the request form and the response text
func (p PatApp) MakeUI() fyne.CanvasObject {

	// list of options for the select GET, POST etc..
	actions := []string{"GET", "POST"}
	sel := widget.NewSelect(actions, func(value string) {
		log.Println("Select set to", value)
	})

	//create a radio group to select HTTP or HTTPS
	radio := widget.NewRadioGroup([]string{"http", "https"}, func(value string) {
		log.Println("Radio set to", value)
	})
	radio.SetSelected("https")
	radio.Horizontal = true

	// this container combines the select drop down and radio button into a container
	// for the right side of a form layout
	method_proto := container.New(layout.NewGridLayout(2), sel, radio)

	//The Host entry field
	host_entry := widget.NewEntry()
	host_entry.SetPlaceHolder("Enter the host name or IP Address... host:port...")

	//The path entry for for setting the URI componenet of the URL
	path_entry := widget.NewEntry()
	path_entry.SetPlaceHolder("Enter the URI componenet with leading / ...")

	path_label := widget.NewLabel("Enter URI:")

	// this container combines a label and a spacer for the left side of line 2 of the form layout
	space_label := container.New(layout.NewGridLayout(3), layout.NewSpacer(), layout.NewSpacer(), path_label)

	//The query string
	query_entry := widget.NewEntry()
	query_entry.SetPlaceHolder("Enter the query string...")

	//prior to here is entry form elements
	//
	//following is the text grid widget that holds the response from the query

	result_text := widget.NewTextGrid()
	result_text.SetText("Now is the time for all food men...")

	main_container := container.NewVBox(
		container.New(layout.NewFormLayout(), method_proto, host_entry,
			space_label, path_entry),
		widget.NewSeparator(),
		result_text,
	)

	return main_container
}

// End of form and Gui creation functions
//
// functions from here are actions required to create the request and display the response

/*
// constructURL takes the form data and constructs a valid URL
// the function also validates that the elements are valid URL components
// returns a validated URL
func constructURL(scheme, host, path, query string) (string, error) {
	if query != "" {
		query = "?" + query
	}
	url := fmt.Sprintf("%s://%s%s%s", scheme, host, path, query)
	//ADD validation of the URl here
	return url, nil
}

// makeRequest takes the http method and url and performs the request
// returns the response struct
func makeRequest(method, url string) (*resty.Response, error) {
	resp, err := conn.Client.R().Get(url)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// responseText takes the request response and extracts response fields and returns a string
// returns a string
//func responseText(resp *resty.Response) string {
//	return fmt.Sprintf("%+v", resp)
//}

*/
