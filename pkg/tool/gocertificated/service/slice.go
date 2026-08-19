package service

func slice(v *[]string) []string {
	if v == nil {
		return nil
	}

	return *v
}
