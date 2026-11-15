package shell_script

func (s *Script) Any(match func(*Call) bool) bool {
	for _, c := range s.Calls {
		if match(c) {
			return true
		}
	}

	return false
}
