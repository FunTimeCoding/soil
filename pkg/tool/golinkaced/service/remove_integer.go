package service

func removeInteger(
	s []int,
	v int,
) []int {
	var result []int

	for _, i := range s {
		if i != v {
			result = append(result, i)
		}
	}

	return result
}
