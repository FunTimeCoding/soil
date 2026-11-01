package service_tester

func addressed(to string) []string {
	if to == "" {
		return nil
	}

	return []string{to}
}
