package module

import "fmt"

func (m *Module) formatVersion() string {
	if m.Latest == m.Version {
		return m.Version
	}

	return fmt.Sprintf("%s -> %s", m.Version, m.Latest)
}
