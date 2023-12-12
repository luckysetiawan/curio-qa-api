package util

func Contains[T comparable](data []T, value T) bool {
	for _, v := range data {
		if v == value {
			return true
		}
	}

	return false
}
