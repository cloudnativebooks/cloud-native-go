package sugar

import "fmt"

func Greeting(prefix string, who ...string) {
	if who == nil {
		fmt.Printf("Nobody to say hi.")
		return
	}

	for _, people := range who {
		fmt.Printf("%s %s\n", prefix, people)
	}
}
