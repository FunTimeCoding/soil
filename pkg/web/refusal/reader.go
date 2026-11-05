package refusal

type Reader func(
	body []byte,
	fallback string,
) string
