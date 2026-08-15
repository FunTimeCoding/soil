package server

func limitOr(
	v *int,
	fallback int,
) int {
	if v == nil || *v <= 0 {
		return fallback
	}

	return *v
}
