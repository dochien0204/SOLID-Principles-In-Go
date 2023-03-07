package ocp

import "fmt"

func Run() {
	fmt.Println("Open/Closed Principle")

	mutationDog := &MutationDogs{}

	fmt.Println(mutationDog.Legs())

	fmt.Println("This dog can bark ? -", mutationDog.IsBark())

}
