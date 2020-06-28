package loop2

import (
	"fmt"
	"time"
)

func foo() {
	var a []int
	a = append(a, 1, 2, 3)
	for _, val := range a {
		go func() {
			fmt.Println(val)
		}()
	}

	time.Sleep(1 * time.Second)
}
