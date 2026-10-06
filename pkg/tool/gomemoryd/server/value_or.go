package server

func valueOr(
	sent *string,
	stored string,
) string {
	if sent == nil {
		return stored
	}

	return *sent
}
