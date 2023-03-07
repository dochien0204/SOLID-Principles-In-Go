package isp

import "fmt"

type ComWithInput struct {
}

func (ci *ComWithInput) Execute() {
	fmt.Println("Com with input execute")
}

func (c *ComWithInput) HandleInput() {
	fmt.Println("Command with input is handling input")
}
