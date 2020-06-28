package _slice

import (
	"testing"
)

var s []string

func setSamples() {
	s = []string{"aaaaa", "bbbbb", "ccccc", "dddddd", "eeeeee", "fffff", "ggggg", "hhhhh", "iiiii", "jjjjj", "kkkkk", "lllll",
		"mmmmm", "nnnnn", "ooooo", "ppppp", "qqqqq", "rrrrr", "sssss", "ttttt", "uuuuu", "vvvvv", "wwwww", "xxxxx", "yyyyy", "zzzzz"}
}

func BenchmarkCopyByIteration(b *testing.B) {
	setSamples()

	for i := 0; i < b.N; i++ {
		CopyByIteration(s)
	}
}

func BenchmarkCopyByBuiltIn(b *testing.B) {
	setSamples()

	for i := 0; i < b.N; i++ {
		CopyByBuiltIn(s)
	}
}
