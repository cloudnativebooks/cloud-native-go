package reflection

import (
	"fmt"
	"reflect"
)

func Law2() {
	var A interface{}
	A = 100

	v := reflect.ValueOf(A)
	B := v.Interface()

	if A == B {
		fmt.Printf("They are same!\n")
	}
}

func Law3Fail() {
	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	v.SetFloat(7.1) // Error: will panic.
	fmt.Printf("x = %f\n", x)
}
