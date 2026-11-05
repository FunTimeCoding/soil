package basic

func reason(
	body []byte,
	fallback string,
) string {
	if p := paragraph(body); p != "" {
		return p
	}

	return fallback
}
