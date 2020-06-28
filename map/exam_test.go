package _map

import "fmt"

// panic: assignment to entry in nil map
func ExampleAddFruit() {
	AddFruit("apple", "red")
}

func ExampleGetScore() {
	s := GetScore("Rainbow")
	fmt.Printf("score is %d\n", s)
	// Output:
	// score is 0
}

func ExampleGetScoreImproved() {
	s := GetScoreImproved("Rainbow")
	fmt.Printf("score is %d\n", s)
	// Output:
	// score is -1
}
