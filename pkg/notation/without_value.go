package notation

func withoutValue(
	a any,
	key string,
) any {
	switch t := a.(type) {
	case map[string]any:
		delete(t, key)

		for k, value := range t {
			t[k] = withoutValue(value, key)
		}

		return t
	case []any:
		for i, value := range t {
			t[i] = withoutValue(value, key)
		}

		return t
	}

	return a
}
