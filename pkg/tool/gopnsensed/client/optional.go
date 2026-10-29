package client

func optional(phrase string) *string {
	if phrase == "" {
		return nil
	}

	return &phrase
}
