package slice

func FindNotEmptyValue[T comparable](slice []T) T {
	var newValue T
	for _, v := range slice {
		if v != newValue {
			return v
		}
	}
	return newValue
}
