package example

func Suppressed() any {
	var result struct { // goanalyze:ignore
		Name string
	}

	return result
}
