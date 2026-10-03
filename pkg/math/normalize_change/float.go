package normalize_change

func Float(
	now float64,
	change float64,
	minimum float64,
	maximum float64,
) float64 {
	if maximum > minimum && now+change > maximum {
		return maximum - now
	}

	if now+change < minimum {
		return (now - minimum) * -1
	}

	return change
}
