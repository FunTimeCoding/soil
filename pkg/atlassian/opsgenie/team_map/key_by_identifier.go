package team_map

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func (m *Map) KeyByIdentifier(identifier string) string {
	t := m.ByIdentifier(identifier)

	if t == nil {
		console.Format("Team not found: %s\n", identifier)

		return constant.OpsgenieNoKey
	}

	return m.KeyByName(t.Name)
}
