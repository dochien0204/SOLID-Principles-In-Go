package ocp

type Dogs struct {
	name string
}

func (d Dogs) Legs() int {
	return 4
}

func (d Dogs) IsBark() bool {
	return true
}

type MutationDogs struct {
	Dogs
}

func (m MutationDogs) Legs() int {
	return 5
}
