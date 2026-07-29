package customer

import (
	"fmt"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func New(v *models.CustomerRequestScheme) *Issue {
	var summary string
	var description string

	for _, e := range v.RequestFieldValues {
		switch e.FieldID {
		case constant.JiraSummaryField:
			summary = fmt.Sprintf("%s", e.Value)
		case constant.JiraDescriptionField:
			description = fmt.Sprintf("%s", e.Value)
		}
	}

	return &Issue{
		Key:    v.IssueKey,
		Status: v.CurrentStatus.Status,
		Title:  summary,
		Body:   description,
		Link:   v.Links.Web,
		Raw:    v,
	}
}
