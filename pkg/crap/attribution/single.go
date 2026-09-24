package attribution

func (m *Matrix) Single() map[string]string {
	result := map[string]string{}

	for key, tests := range m.Covers {
		if len(tests) == 1 {
			result[key] = tests[0]
		}
	}

	return result
}
