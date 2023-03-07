package isp

import "fmt"

func Run() {
	fmt.Println()
	fmt.Println("Interface Segregation Principle")
	com := &ComWithInput{}

	com.HandleInput()
	com.Execute()
}
