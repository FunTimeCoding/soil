package goprocess

func each(
	names []string,
	f func(string) error,
) error {
	for _, n := range names {
		if e := f(n); e != nil {
			return e
		}
	}

	return nil
}
