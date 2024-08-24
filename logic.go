package main

import (
	"net"
	"os/exec"
	"strings"
)

func setup(con_name string, iface string, ssid string, psk string, is_hidden bool, auto_con bool, con_now bool) (labelVal string) {
	// Setup the hidden connection status for nmcli
	hidden := "no"

	if is_hidden {
		hidden = "yes"
	}

	// Setup the auto connection status for nmcli
	auto := "no"

	if auto_con {
		auto = "yes"
	}

	// Create the connection
	create := exec.Command("flatpak-spawn", "--host", "nmcli", "c", "add", "type", "wifi", "con-name", con_name, "ifname", iface, "ssid", ssid)
	_, err := create.Output()
	if err != nil {
		labelVal = (string("There was an error creating the wifi network connection ") + err.Error())
		return
	}

	// Add the security to the wifi network created
	sec := exec.Command("flatpak-spawn", "--host", "nmcli", "con", "modify", con_name, "wifi-sec.key-mgmt", "wpa-psk", "wifi-sec.psk", psk, "wifi.hidden", hidden, "connection.autoconnect", auto)
	_, err = sec.Output()
	if err != nil {
		labelVal = "There was an error setting up the security and auto-connect status of the new connection" + err.Error()
		return
	}

	// Connect to the new network
	if con_now {
		con := exec.Command("flatpak-spawn", "--host", "nmcli", "con", "up", con_name)
		_, err = con.Output()
		if err != nil {
			labelVal = "There was an error connecting to " + con_name + ":\n " + err.Error()
			return
		}

		labelVal = "Successfully created and connected to the network!"
	} else {
		labelVal = "Successfully created the nework!"
	}

	return
}

func remove(con_name string) (labelVal string) {
	remove := exec.Command("flatpak-spawn", "--host", "nmcli", "con", "delete", "id", con_name)
	out, err := remove.Output()
	if err != nil {
		labelVal = "There was an error removing " + con_name
		return
	}

	labelVal = string(out)

	return
}

func getCurrentlySetupCons() (cons []string) {
	active_cons := exec.Command("flatpak-spawn", "--host", "nmcli", "con", "show")
	out, err := active_cons.Output()
	if err != nil {
		cons = append(cons, "No network connections have been setup...")
		return
	}

	cons = strings.Split(string(out), "\n")
	return
}

// Helper functions
func getInterfaceNames() (ifaces []string) {
	ifs, err := net.Interfaces()
	if err != nil {
		ifaces = append(ifaces, "The interface names could not be loaded!")
		return
	}

	for _, i := range ifs {
		ifaces = append(ifaces, i.Name)
	}

	return
}
