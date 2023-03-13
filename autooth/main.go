package main

import (
	"fmt"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/device"
	_ "github.com/muka/go-bluetooth/bluez/profile/device"
)

func main() {

	// THIS CODE WORKS
	// Get the adapter for your Bluetooth dev
	adapter, err := api.GetDefaultAdapter()
	if err != nil {
		fmt.Println("Failed to get adapter:", err)
		return
	}

	// Start the adapter
	err = adapter.StartDiscovery()
	if err != nil {
		fmt.Println("Failed to start discovery:", err)
		return
	}

	// Get adapter list of available Bluetooth devices
	devices, err := adapter.GetDevices()
	if err != nil {
		fmt.Println("Failed to get devices:", err)
		return
	}

	for _, dev := range devices {
		fmt.Println(dev.Properties.Name)
	}

	// Check if your Bluetooth speaker is in the list of available devices
	var speaker *device.Device1
	for _, dev := range devices {
		if dev.Properties.Name == "T2" {
			speaker = dev
			break
		}
	}

	// Connect to your Bluetooth speaker
	err = speaker.Connect()
	if err != nil {
		fmt.Println("Failed to connect to speaker:", err)
		return
	}

	// Play audio through your Bluetooth speaker
	fmt.Println("Connected to speaker:", speaker.Properties.Name)
}
