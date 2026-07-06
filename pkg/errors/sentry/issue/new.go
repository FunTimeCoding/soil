package issue

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/monitor/item/constant"
)

func New(v *response.Issue) *Issue {
	return &Issue{
		MonitorIdentifier: constant.GoSentry.StringIdentifier(v.ID),
		Project:           v.Project.Name,
		Type:              v.Type,
		Title:             v.Title,
		Link:              v.Permalink,
		Create:            v.FirstSeen,
		Raw:               v,
	}
}
