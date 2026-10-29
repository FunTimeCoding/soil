package service

func one(
	values []string,
	e error,
) ([]string, []error) {
	if e == nil {
		return values, nil
	}

	return values, []error{e}
}
