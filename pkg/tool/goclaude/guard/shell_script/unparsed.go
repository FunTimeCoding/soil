package shell_script

func (s *Script) Unparsed() bool {
	if !s.Parsed {
		return true
	}

	for _, n := range s.Scripts {
		if !n.Parsed {
			return true
		}
	}

	return false
}
