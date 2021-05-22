package utils

func ReverseMap(src map[int]string) (dst map[string]int) {

	if src == nil {
		return
	}

	dst = map[string]int{}

	for key, value := range src {

		if _, found := dst[value]; found {
			panic("key is already present")
		}

		dst[value] = key
	}

	return
}
