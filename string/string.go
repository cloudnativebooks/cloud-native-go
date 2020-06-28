package _string

import "fmt"

func StringDoubleQuotationMarks() {
	s := "Hi, \nthis is \"RainbowMango\"."
	fmt.Println(s)
}

func StringBackslash() {
	s := `Hi, 
this is "RainbowMango".`
	fmt.Println(s)
}

func StringConnectSample1() string {
	var s string
	s = s + "a" + "b"
	return s
}

func StringConnectSample2() string {
	var s string
	s = s + "a"
	s = s + "b"
	return s
}

func StringToByte() {
	b := []byte{'H', 'e', 'l', 'l', 'o'}
	s := string(b)
	fmt.Println(s) // Hello
}

func ByteToString() {
	s := "Hello"
	b := []byte(s)
	fmt.Println(b) // [72 101 108 108 111]
}

func StringIteration() {
	s := "中国"
	for index, value := range s {
		fmt.Printf("index: %d, value: %c\n", index, value)
	}
}
