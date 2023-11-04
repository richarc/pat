package gui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-resty/resty/v2"
	"github.com/richarc/pat/conn"
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

	// the text widget that we will write the
	response_text := func() *widget.TextGrid {
		grid := widget.NewTextGrid()
		grid.SetText("Response text goes here")
		return grid
	}()

	// list of options for the selct GET, POST etc..
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

	//The Host entry field
	host_entry := widget.NewEntry()
	host_entry.SetPlaceHolder("Enter the host name or IP Address...")

	//The Host entry field
	uri_entry := widget.NewEntry()
	uri_entry.SetPlaceHolder("Enter the URI and query string...")

	request_form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Select HTTP Method:", Widget: sel},
			{Text: "Select HTTP or HTTPS:", Widget: radio},
			{Text: "Enter the Host Address:", Widget: host_entry},
			{Text: "Enetr the URI and Query Componenet:", Widget: uri_entry},
		},
		OnSubmit: func() {
			// construct the URL and check for errors
			url, err := constructURL(radio.Selected, host_entry.Text, uri_entry.Text, "")
			if err != nil {
				log.Println("error in URL")
			}
			// make the request
			resp, err := makeRequest(sel.Selected, url)
			if err != nil {
				log.Println(err)
			}
			log.Println(resp)
			//display the respons in the responsetext widget
			response_text.SetText(fmt.Sprintf("%+v", resp))
		},
		OnCancel: func() {
			log.Println("Cancel Clicked")
		},
		SubmitText: "Run",
		CancelText: "Reset",
	}

	main_container := container.NewVBox(request_form, widget.NewSeparator(), response_text)
	return main_container
}

// End of form and Gui creation functions
//
// functions from here are actions required to create the request and display the response

// constructURL takes the Dorm data and constructs a valid URL
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
