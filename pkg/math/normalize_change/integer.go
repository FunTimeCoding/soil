package normalize_change

func Integer(
	now int,
	change int,
	minimum int,
	maximum int,
) int {
	if maximum > minimum && now+change > maximum {
		return maximum - now
	}

	if now+change < minimum {
		return (now - minimum) * -1
	}

	return change
}
