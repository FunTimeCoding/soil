package attribution

func (m *Matrix) Add(
	test string,
	functionKeys []string,
) {
	m.Tests = append(m.Tests, test)

	for _, k := range functionKeys {
		m.Covers[k] = append(m.Covers[k], test)
	}
}
