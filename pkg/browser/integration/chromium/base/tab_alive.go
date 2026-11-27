package base

func (s *Stack) TabAlive(identifier string) bool {
	s.T.Helper()

	for _, t := range s.tabs() {
		if t.Identifier == identifier {
			return true
		}
	}

	return false
}
