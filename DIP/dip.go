package dip

import "fmt"

func Run() {
	fmt.Println()
	fmt.Println("Dependency Inversion Principle")
	b1 := &Book{
		1,
		"Book 1",
	}

	b2 := &Book{
		2,
		"Book 2",
	}

	books := []Book{}
	user := &User{
		Books: books,
	}
	user.AddBookForUser(*b1)
	user.AddBookForUser(*b2)
	fmt.Println(user.Books)
	user.GetAllBook()
}
