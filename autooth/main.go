package main

import "fmt"

func main() {
	var tooth Autooth
	err := tooth.initialize()
	if err != nil {
		panic(err)
	}

	err = tooth.ConnectToDevice("T2")
	if err != nil {
		fmt.Println()
	}

	/*err = tooth.Disconnect()*/
}
