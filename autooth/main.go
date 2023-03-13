package main

import "fmt"

func main() {
	var tooth Autooth
	err := tooth.initialize()
	if err != nil {
		fmt.Println(err)
	}

	err = tooth.ConnectToDevice("a")
	if err != nil {
		fmt.Println(err)
	}

}
