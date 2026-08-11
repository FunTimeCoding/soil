package format

func pairMark(
	declared bool,
	served bool,
) string {
	if declared && served {
		return "+"
	}

	if declared || served {
		return "~"
	}

	return "-"
}
