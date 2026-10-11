package module

func (m *Module) HasConcerns() bool {
	return len(m.concern) > 0
}
