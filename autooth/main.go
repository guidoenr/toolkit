package main

func main() {
	var tooth Autooth
	err := tooth.initialize()
	if err != nil {
		panic(err)
	}

	err = tooth.ConnectToDevice("T2")
	if err != nil {
		panic(err)
	}

}
