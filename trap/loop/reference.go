package loop

import "fmt"

func ReferenceLoopIteratorVariable() {
	var out []*int
	for i := 0; i < 3; i++ {
		out = append(out, &i)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("cap:", cap(out))
}
