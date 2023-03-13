package main

import (
	"errors"
	"fmt"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/bluez/profile/device"
)

// Autooth (stands for auto-bluetooth) has the entire logic to connect to your desired bluetooth devices
type Autooth struct {
	Adapter *adapter.Adapter1
	Devices []*device.Device1
}

func (a *Autooth) initialize() error {
	// get the default adapter
	defaultAdapter, err := api.GetDefaultAdapter()
	if err != nil {
		msg := fmt.Sprintf("getting adapter: %v", err)
		return errors.New(msg)
	}

	// set the adapter
	a.Adapter = defaultAdapter

	return nil
}

// ScanDevices scan the devices and set it
func (a *Autooth) ScanDevices() error {

	// start the discovering process
	err := a.Adapter.StartDiscovery()
	if err != nil {
		msg := fmt.Sprintf("discovering: %v", err)
		return errors.New(msg)
	}

	// get adapter list of available Bluetooth devices
	devices, err := a.Adapter.GetDevices()
	if err != nil {
		msg := fmt.Sprintf("getting devices: %v", err)
		return errors.New(msg)
	}

	// set the devices
	a.Devices = devices

	return nil
}

// ConnectToDevice scan the devices and set it
func (a *Autooth) ConnectToDevice(deviceName string) error {
	// Check if your Bluetooth speaker is in the list of available devices
	var speaker *device.Device1
	for _, dev := range a.Devices {
		fmt.Println(dev.Properties.Name)
		if dev.Properties.Name == deviceName {
			speaker = dev
			break
		}
	}

	if speaker == nil {
		msg := fmt.Sprintf("device '%s' not found", deviceName)
		return errors.New(msg)
	}

	// Connect to your Bluetooth speaker
	err := speaker.Connect()
	if err != nil {
		msg := fmt.Sprintf("connecting to speaker: %v", err)
		return errors.New(msg)
	}

	return nil
}