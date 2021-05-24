package utils

import "errors"

func ReverseMap(src map[int]string) (dst map[string]int, err error) {

	if src == nil {
		err = errors.New("src is nil")
		return
	}

	if len(src) == 0 {
		return
	}

	dst = make(map[string]int, len(src))

	for key, value := range src {

		if _, found := dst[value]; found {
			err = errors.New("key is already present")
			return
		}

		dst[value] = key
	}

	return
}
