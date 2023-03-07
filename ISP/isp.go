package isp

import "fmt"

func Run() {
	fmt.Println()
	fmt.Println("Interface Segregation Principle")

	//command with input
	var iCommandWithInput ICommandWithInput = &ComWithInput{}
	iCommandWithInput.HandleInput()
	iCommandWithInput.Execute()

	var iCommandWithoutInput ICommand = &Com{}
	iCommandWithoutInput.Execute()
}
