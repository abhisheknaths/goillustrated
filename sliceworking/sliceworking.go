package sliceworking

func AppendWithCapacity(loopSize int) {
	a := make([]int, 0, loopSize)
	for i := 0; i < loopSize; i++ {
		a = append(a, i)
	}
}

func AppendWithoutCapacity(loopSize int) {
	var a []int
	for i := 0; i < loopSize; i++ {
		a = append(a, i)
	}
}
