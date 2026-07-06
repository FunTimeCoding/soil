package issue

import "github.com/funtimecoding/soil/pkg/atlassian/jira/field_map"

func (i *Issue) FieldMap() *field_map.Map {
	return i.fieldMap
}
