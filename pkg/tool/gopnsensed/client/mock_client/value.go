package mock_client

func value(v *string) string {
	if v == nil {
		return ""
	}

	return *v
}
