package srp

type User struct {
	id    string
	Name  string
	Phone string
}

func (u *User) GetSender() *User {
	return &User{
		id:    "U01",
		Name:  "Chien sender",
		Phone: "0123456789",
	}
}

func (u *User) GetReceiver() *User {
	return &User{
		id:    "U02",
		Name:  "Chien receiver",
		Phone: "0123456789",
	}
}
