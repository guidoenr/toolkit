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
	Adapter         *adapter.Adapter1
	Devices         []*device.Device1
	ConnectedDevice *device.Device1
}

// initialize the autooth client
func (a *Autooth) initialize() error {
	// get the default adapter
	defaultAdapter, err := api.GetDefaultAdapter()
	if err != nil {
		msg := fmt.Sprintf("getting adapter: %v", err)
		return errors.New(msg)
	}

	// set the adapter
	a.Adapter = defaultAdapter

	// ------------------------------------- PROPERTIES
	// setting the properties
	err = a.Adapter.SetProperty("Powered", true)
	if err != nil {
		return err
	}

	// set discoverable
	err = a.Adapter.SetProperty("Discoverable", true)
	if err != nil {
		return err
	}

	err = a.Adapter.SetProperty("DiscoverableTimeout", uint32(0))
	if err != nil {
		return err
	}

	// set discoverable timeout
	err = a.Adapter.SetDiscoverableTimeout(120)
	if err != nil {
		return err
	}

	return nil
}

// ScanDevices scan the devices and set it
func (a *Autooth) ScanDevices() error {
	// start the discovering process
	err := a.Adapter.StartDiscovery()
	if err != nil {
		msg := fmt.Sprintf("discovering: %v \n", err)
		return errors.New(msg)
	}

	// get adapter list of available Bluetooth devices
	devices, err := a.Adapter.GetDevices()
	if err != nil {
		msg := fmt.Sprintf("getting devices: %v \n", err)
		return errors.New(msg)
	}

	// set the devices
	a.Devices = devices
	fmt.Printf("%d devices available to connect \n", len(a.Devices))
	for _, d := range a.Devices {
		fmt.Println(d.Properties.Name)
	}

	return nil
}

// ConnectToDevice scan the devices and set it
func (a *Autooth) ConnectToDevice(deviceName string) error {
	var deviceToConnect *device.Device1

	// scan the devices
	err := a.ScanDevices()
	if err != nil {
		return err
	}

	// check all the devices and connect to the deviceName
	for _, dev := range a.Devices {
		fmt.Println(dev.Properties.Name)
		if dev.Properties.Name == deviceName {
			deviceToConnect = dev
			break
		}
	}

	// if the speaker doesn't exists
	if deviceToConnect == nil {
		msg := fmt.Sprintf("device '%s' not found \n", deviceName)
		return errors.New(msg)
	}

	// get the connected status
	connected, err := deviceToConnect.GetConnected()
	if err != nil {
		msg := fmt.Sprintf("getConnected: %v", err)
		return errors.New(msg)
	}

	// already connected
	if connected {
		fmt.Printf("already connected to device: %s \n", deviceToConnect.Properties.Name)
		a.ConnectedDevice = deviceToConnect
		return nil
	}

	// connect the bluetooth speaker
	err = deviceToConnect.Connect()
	if err != nil {
		msg := fmt.Sprintf("connecting to speaker: %v \n", err)
		return errors.New(msg)
	}

	// set trusted
	err = deviceToConnect.SetTrusted(true)
	if err != nil {
		msg := fmt.Sprintf("setting trusted: %v", err)
		return errors.New(msg)
	}

	// set paired
	/*	err = deviceToConnect.SetPaired(true)
		if err != nil {
			msg := fmt.Sprintf("setting paired: %v", err)
			return errors.New(msg)
		}*/

	// set the connected device
	a.ConnectedDevice = deviceToConnect

	return nil
}

// Disconnect disconnects the device
func (a *Autooth) Disconnect() error {
	// disconnect from the device
	err := a.ConnectedDevice.Disconnect()
	if err != nil {
		msg := fmt.Sprintf("disconnecting from '%v': %v \n", a.ConnectedDevice, err)
		return errors.New(msg)
	}
	fmt.Printf("disconnected from device: %s \n", a.ConnectedDevice.Properties.Name)
	return nil
}

/*
func (a *Autooth) Ping() {
	fmt.Println("[Autooth status]")
	fmt.Printf("connected device: %s \n", a.ConnectedDevice.Properties.Name)
	fmt.Printf("adapter name: %s \n", a.Adapter.Properties.Name)
}
*/
