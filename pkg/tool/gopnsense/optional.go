package gopnsense

func optional(v string) *string {
	if v == "" {
		return nil
	}

	return &v
}
