package utils

func SplitSlice(src []int, chunkSize int) (dst [][]int) {

	if src == nil || chunkSize <= 0 {
		return
	}

	dst = [][]int{}

	sliceLen := len(src)

	for i := 0; i < sliceLen/chunkSize; i++ {

		begin := 0 + chunkSize*i
		end := chunkSize + chunkSize*i

		dst = append(dst, src[begin:end])
	}

	left := sliceLen % chunkSize

	if left > 0 {

		dst = append(dst, src[sliceLen-left:])
	}

	return
}
