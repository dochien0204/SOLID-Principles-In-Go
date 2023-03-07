package isp

import "fmt"

type Com struct {
	HasInput bool
}

func (c *Com) Execute() {
	fmt.Println("Command Executed")
}
