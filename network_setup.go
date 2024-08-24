package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Network Connection")
	w.Resize(fyne.NewSize(1024.0, 768.0))
	ifaceNames := getInterfaceNames()

	// setup the form

	conNameText := widget.NewEntry()
	conNameText.SetPlaceHolder("Connection Display Name")
	ssidText := widget.NewEntry()
	ssidText.SetPlaceHolder("SSID (Network Name)")

	iface_selected := string("")
	ifaces := widget.NewSelect(
		ifaceNames, func(opt string) {
			iface_selected = opt
		})

	psk := widget.NewPasswordEntry()
	psk.SetPlaceHolder("WiFi Password")

	is_hidden := false
	isHiddenCheck := widget.NewCheck("Hidden Network", func(hidden bool) {
		is_hidden = hidden
	})

	auto_con := false
	autoConCheck := widget.NewCheck("Auto-connect?", func(auto bool) {
		auto_con = auto
	})

	con_now := false
	conNowCheck := widget.NewCheck("Connect on Creation?", func(con bool) {
		con_now = con
	})

	cur_setup_connectionsLabel := widget.NewLabel("")
	cons := getCurrentlySetupCons()
	for con := range cons {
		cur_setup_connectionsLabel.Text += string(cons[con] + "\n")
	}
	status := widget.NewLabel("Status...")

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

	w.SetContent(form)

	w.ShowAndRun()
}
