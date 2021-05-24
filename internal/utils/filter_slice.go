package utils

import "errors"

func FilterSlice(src []int, keys []int) (dst []int, err error) {

	if src == nil {
		err = errors.New("src is nil")
		return
	}

	if keys == nil {
		err = errors.New("keys is nil")
		return
	}

	if len(src) == 0 {
		return
	}

	var filter = map[int]struct{}{}

	for _, key := range keys {
		filter[key] = struct{}{}
	}

	for i := 0; i < len(src); i++ {

		if _, found := filter[src[i]]; !found {
			dst = append(dst, src[i])
		}
	}

	return
}
