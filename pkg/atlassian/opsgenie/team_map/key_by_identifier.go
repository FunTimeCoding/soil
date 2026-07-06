package team_map

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/constant"
)

func (m *Map) KeyByIdentifier(identifier string) string {
	t := m.ByIdentifier(identifier)

	if t == nil {
		fmt.Printf("Team not found: %s\n", identifier)

		return constant.NoKey
	}

	return m.KeyByName(t.Name)
}
