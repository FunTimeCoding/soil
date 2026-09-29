package strings

func MustToFloat(s string) float64 {
	result, e := ParseFloat(s)

	if e != nil {
		panic(e)
	}

	return result
}
