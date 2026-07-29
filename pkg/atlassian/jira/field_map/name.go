package field_map

import "github.com/funtimecoding/soil/pkg/atlassian/constant"

func (m *Map) Name(identifier string) string {
	f := m.ByIdentifier(identifier)

	if f == nil {
		return constant.JiraUnknown
	}

	return f.Name
}
