package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/field_map"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/option"
	"github.com/funtimecoding/soil/pkg/face"
	"time"
)

type Issue struct {
	MonitorIdentifier string
	Key               string
	Summary           string
	Description       string
	Initials          string
	Status            string
	Priority          string
	Type              string
	Create            *time.Time
	Due               *time.Time
	Link              string
	Labels            []string
	concerns          []string
	score             float64
	commentNameFilter []string
	shortStatus       face.StringAlias
	scoreColor        face.SprintFunction
	ageColor          face.SprintFunction
	fieldMap          *field_map.Map
	option            *option.Issue
	Raw               *jira.Issue
}
