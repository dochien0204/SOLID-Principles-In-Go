package srp

import "fmt"

type Message struct {
	Body string
}

func (message *Message) Create() string {

	message = &Message{
		Body: "Hello",
	}

	rs := message.Body

	fmt.Println("Message created!")

	return rs
}
