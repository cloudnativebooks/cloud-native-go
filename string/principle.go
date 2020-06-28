package _string

import "fmt"

func GetUnicode() {
	s := "中国"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func GetAscii() {
	s := "hi"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
