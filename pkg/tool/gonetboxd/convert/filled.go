package convert

func filled(v string) *string {
	if v == "" {
		return nil
	}

	return &v
}
