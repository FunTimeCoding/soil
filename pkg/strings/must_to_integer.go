package strings

func MustToInteger(s string) int {
	result, e := ParseInteger(s)

	if e != nil {
		panic(e)
	}

	return result
}
