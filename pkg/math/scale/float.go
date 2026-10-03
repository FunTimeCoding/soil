package scale

func Float(
	from float64,
	to float64,
	factor float64,
) float64 {
	if factor < 0 {
		factor = 0
	} else if factor > 1 {
		factor = 1
	}

	return from + (to-from)*factor
}
