package server

func text(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}
