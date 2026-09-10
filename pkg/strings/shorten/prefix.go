package shorten

func Prefix(value string, count int) string {
	r := []rune(value)

	if len(r) <= count {
		return value
	}

	return string(r[:count])
}
