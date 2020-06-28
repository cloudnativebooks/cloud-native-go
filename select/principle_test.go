package _select

func ExampleSelectGo() {
	var cas0 = [2]scase{
		{c: "1"},
		{c: "2"},
	}

	var order0 = [4]uint16{}

	SelectGo(&cas0[0], &order0[0], 2)
	// Output:
	// len(cas1)=65536, cap(cas1)=65536
	// len(order1)=131072, cap(order1)=131072
	// len(scases)=2, cap(scases)=2
	// len(pollorder)=2, cap(pollorder)=2
	// len(lockorder)=2, cap(lockorder)=2
}

func ExampleWriteNilChannle() {
	WriteNilChannle()
	// Output:
	// default
}
