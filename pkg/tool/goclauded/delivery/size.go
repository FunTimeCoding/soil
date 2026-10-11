package delivery

func size(lines []string) int {
	var result int

	for _, l := range lines {
		result += lineSize(l)
	}

	return result
}
