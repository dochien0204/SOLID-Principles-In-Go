package srp

import "fmt"

func Run() {

	fmt.Println("Single resibility principle")
	message := &Message{}

	messagePayload := message.Create()

	user := &User{}

	sender := &Sender{
		SenderPerson:  user.GetSender(),
		ReceivePerson: user.GetReceiver(),
		Message:       messagePayload,
	}

	sender.SendMessage()
	fmt.Println()
}
