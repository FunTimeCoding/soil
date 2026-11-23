package convert

func text(v *string) string {
	if v == nil {
		return ""
	}

	return *v
}
