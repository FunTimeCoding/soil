package model_context

func given(v *[]string) []string {
	if v == nil {
		return nil
	}

	return *v
}
