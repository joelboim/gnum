package infra

func GetPointerValue[T any](value *T, defaultValue T) T {
	if value != nil {
		return *value
	}

	return defaultValue
}
