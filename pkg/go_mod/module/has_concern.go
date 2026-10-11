package module

import "slices"

func (m *Module) HasConcern(s string) bool {
	return slices.Contains(m.concern, s)
}
