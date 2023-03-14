package main

import "fmt"

func main() {
	var tooth Autooth
	err := tooth.initialize()
	if err != nil {
		panic(err)
	}

	// ADDRESS :
	err = tooth.ConnectToDevice("JBL-guido-c5")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("name: %s", tooth.ConnectedDevice.Properties.Name)

	/*	tooth.Ping()
	 */
	/*err = tooth.Disconnect()*/
}
