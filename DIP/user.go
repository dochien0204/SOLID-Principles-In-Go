package dip

import "fmt"

type User struct {
	Books []Book
}

func (u *User) GetAllBook() {
	for i, book := range u.Books {
		var iBook IBook = &book
		fmt.Printf("Information of Book %d\n", i)
		iBook.GetInformationBook()
	}
}

func (u *User) AddBookForUser(book Book) bool {
	u.Books = append(u.Books, book)
	return true
}
