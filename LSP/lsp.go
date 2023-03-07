package lsp

import "fmt"

func Run() {

	fmt.Println()
	fmt.Println("Liskov Substitution Principle")
	rectangle := &Square{}

	rectangle.SetHeight(100)

	fmt.Printf("Area = %d\n", rectangle.GetArea())
}
