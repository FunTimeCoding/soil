package option

import "github.com/funtimecoding/soil/pkg/atlassian/jira/field_map"

type Issue struct {
	Locator       string
	User          string
	WatchedIssues []string
	WatchedLoaded bool
	FieldMap      *field_map.Map
	Verbose       bool
	ClosedStatus  []string
}
