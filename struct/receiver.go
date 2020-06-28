package _struct

import "fmt"

type Student struct {
	Name string
}

// 作用于Student的拷贝对象，修改不会反映到原对象
func (s Student) SetName(name string) {
	s.Name = name
}

// 作用于Student的指针对象，修改会反映到原对象
func (s *Student) UpdateName(name string) {
	s.Name = name
}

func Receiver() {
	s := Student{}
	s.SetName("Rainbow")
	fmt.Printf("Name: %s\n", s.Name) // empty
	s.UpdateName("Rainbow")
	fmt.Printf("Name: %s\n", s.Name) // Rainbow
}
