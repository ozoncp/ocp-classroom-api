package utils

func SplitSliceToChunks(src []int, chunkSize int) (dst [][]int) {

	if chunkSize <= 0 {
		return
	}

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

func ExampleSplitSliceToChunks() {

}
