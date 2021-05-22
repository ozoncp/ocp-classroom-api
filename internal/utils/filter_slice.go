package utils

func FilterSlice(src []int, keys []int) (dst []int) {

	if src == nil || keys == nil {
		return
	}

	var filter = map[int]struct{}{}

	for _, key := range keys {
		filter[key] = struct{}{}
	}

	dst = []int{}

	for i := 0; i < len(src); i++ {

		if _, found := filter[src[i]]; !found {
			dst = append(dst, src[i])
		}
	}

	return
}
