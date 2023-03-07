package srp

import "fmt"

type Sender struct {
	SenderPerson  *User
	ReceivePerson *User
	Message       string
}

func (s *Sender) SendMessage() bool {
	fmt.Println("Send a message")
	fmt.Println("Sender Person:", s.SenderPerson.Name)
	fmt.Println("Receive Person:", s.SenderPerson.Name)
	fmt.Println("Message:", s.Message)
	return true
}
