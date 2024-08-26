package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create the application
	a := app.New()
	// Create the main window
	w := a.NewWindow("Network Connection")
	// Set the window size
	w.Resize(fyne.NewSize(1024.0, 768.0))

	// Get the network interface names
	ifaceNames := getInterfaceNames()

	// setup the form widgets
	conNameText := widget.NewEntry()
	conNameText.SetPlaceHolder("Connection Display Name")
	ssidText := widget.NewEntry()
	ssidText.SetPlaceHolder("SSID (Network Name)")

	// Create a selected variable for the interface names
	iface_selected := string("")
	// Load the select list with the avaliable interface names
	ifaces := widget.NewSelect(
		ifaceNames, func(opt string) {
			iface_selected = opt
		})

	// Setup the pre-shared key password entry widget
	psk := widget.NewPasswordEntry()
	psk.SetPlaceHolder("WiFi Password")

	// Setup the hidden network checkbox and its functionality.
	is_hidden := false
	isHiddenCheck := widget.NewCheck("Hidden Network", func(hidden bool) {
		is_hidden = hidden
	})

	// Setup the auto connect checkbox and its functionality.
	auto_con := false
	autoConCheck := widget.NewCheck("Auto-connect?", func(auto bool) {
		auto_con = auto
	})

	// Setup the connection on creation checkbox and its functionality
	con_now := false
	conNowCheck := widget.NewCheck("Connect on Creation?", func(con bool) {
		con_now = con
	})

	// Create the label to hold the currently created network connections
	cur_setup_connectionsLabel := widget.NewLabel("")

	// Get the currently created network connections and populate the label text
	cons := getCurrentlySetupCons()
	for con := range cons {
		cur_setup_connectionsLabel.Text += string(cons[con] + "\n")
	}

	// Create the action status label
	status := widget.NewLabel("Status...")

	// Create and setup the form with the widgets from above.
	form := widget.NewForm(
		widget.NewFormItem("Connection Name", conNameText),
		widget.NewFormItem("SSID", ssidText),
		widget.NewFormItem("Interfaces", ifaces),
		widget.NewFormItem("WiFi Password", psk),
		widget.NewFormItem("", isHiddenCheck),
		widget.NewFormItem("", autoConCheck),
		widget.NewFormItem("", conNowCheck),
		widget.NewFormItem("", widget.NewButton("Create Network", func() {
			status.Text = setup(conNameText.Text, iface_selected, ssidText.Text, psk.Text, is_hidden, auto_con, con_now)
			status.Refresh()
			cons = getCurrentlySetupCons()
			cur_setup_connectionsLabel.Text = ""
			for con := range cons {
				cur_setup_connectionsLabel.Text += string(cons[con] + "\n")
			}
			cur_setup_connectionsLabel.Refresh()
		})),
		widget.NewFormItem("", widget.NewButton("Remove Connection", func() {
			status.Text = remove(conNameText.Text)
			status.Refresh()
			cons = getCurrentlySetupCons()
			cur_setup_connectionsLabel.Text = ""
			for con := range cons {
				cur_setup_connectionsLabel.Text += string(cons[con] + "\n")
			}
			cur_setup_connectionsLabel.Refresh()
		})),
		widget.NewFormItem("", status),
		widget.NewFormItem("Current Setup Connections", cur_setup_connectionsLabel))

	// Set the main window content to the form
	w.SetContent(form)
	// Show and run the window/application
	w.ShowAndRun()
}
