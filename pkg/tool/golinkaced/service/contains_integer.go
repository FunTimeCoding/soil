package service

func containsInteger(
	s []int,
	v int,
) bool {
	for _, i := range s {
		if i == v {
			return true
		}
	}

	return false
}
