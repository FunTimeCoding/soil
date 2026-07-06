package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/option"
	"github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"strings"
	"time"
)

func New(
	v *jira.Issue,
	o *option.Issue,
) *Issue {
	result := Stub()
	result.MonitorIdentifier = constant.GoJira.StringIdentifier(
		v.Key,
	)
	result.Key = v.Key
	result.Summary = strings.TrimSpace(v.Fields.Summary)
	result.Description = strings.TrimSpace(v.Fields.Description)
	result.Initials = initialsField(v)
	result.Status = statusField(v)
	result.Priority = priorityField(v)
	result.Type = v.Fields.Type.Name
	result.Create = new(time.Time(v.Fields.Created))
	result.Due = new(time.Time(v.Fields.Duedate))
	result.Link = BuildLink(o.Locator, v.Key)
	result.Labels = v.Fields.Labels
	result.fieldMap = o.FieldMap
	result.option = o
	result.Raw = v

	return result
}
