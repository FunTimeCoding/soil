package unit

func notNil(s []string) []string {
	if s == nil {
		return []string{}
	}

	return s
}
