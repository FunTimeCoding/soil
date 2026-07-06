package issue

import "github.com/funtimecoding/soil/pkg/atlassian/jira/field_map"

func (i *Issue) SetFieldMap(m *field_map.Map) {
	i.fieldMap = m
}
