package convert

func optional(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}
