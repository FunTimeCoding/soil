package goquery

func label(title string) string {
	if title == "" {
		return "(before first heading)"
	}

	return title
}
