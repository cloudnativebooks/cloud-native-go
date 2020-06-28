package _string

import (
	"fmt"
	"strings"
)

func Contains() {
	s := "Hello World!"
	sub := "World"
	if strings.Contains(s, sub) {
		fmt.Printf("%s contains %s\n", s, sub)
	}
}

func Trim() {
	s := "abba"
	cutset := "ba"
	s = strings.Trim(s, cutset)
	fmt.Printf("Trimed: %s", s)
}

func Replace() {
	s := "Hello, Hello, Hello"
	old := "Hello"
	new := "hello"
	s = strings.Replace(s, old, new, 2)
	fmt.Printf("Replaced: %s\n", s)
}
