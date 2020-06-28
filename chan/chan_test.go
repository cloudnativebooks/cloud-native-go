package _chan

func ExampleReadChanReturnValue() {
	ReadChanReturnValue()
	// Output:
	// length of channel is: 2
	// second return value before channel closed is: true
	// second return value after channel(have data) closed is: true
	// second return value after channel(no data) closed is: false
}
