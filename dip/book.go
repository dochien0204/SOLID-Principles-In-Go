package dip

import "fmt"

type Book struct {
	Id   int
	Name string
}

func (b *Book) GetInformationBook() {
	fmt.Printf("ID: %d\n", b.Id)
	fmt.Printf("Name: %s\n", b.Name)
}
