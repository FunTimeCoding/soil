package store_tester

func (s *Tester) Reveal(
	path string,
	key string,
) (string, error) {
	if key == s.tamper {
		return "tampered", nil
	}

	return s.stored[path][key], nil
}
