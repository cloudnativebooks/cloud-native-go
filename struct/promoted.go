package _struct

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) SetName(name string) {
	a.Name = name
}

type Cat struct {
	Animal
}

func EmbeddedFoo() {
	c := Cat{}

	c.SetName("A")
	fmt.Printf("Name: %s\n", c.Name) // A

	c.Animal.SetName("a")
	fmt.Printf("Name: %s\n", c.Name) // a

	c.Animal.Name = "B"
	fmt.Printf("Name: %s\n", c.Name) // B

	c.Name = "b"
	fmt.Printf("Name: %s\n", c.Name) // b
}
