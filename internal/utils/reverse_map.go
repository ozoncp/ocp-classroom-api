package utils

func ReverseMap(src map[int]string) (dst map[string]int) {

	if src == nil {
		return
	}

	dst = make(map[string]int)

	for key, value := range src {
		dst[value] = key
	}

	return
}
