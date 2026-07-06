package alert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert/option"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team"
	"github.com/funtimecoding/soil/pkg/monitor/item/constant"
)

func (a *Alert) enrich(
	p *option.Alert,
	t *team.Team,
) *Alert {
	a.MonitorIdentifier = constant.GoGenie.IntegerIdentifier(
		a.SmallIdentifier,
	)
	a.TeamKey = p.Team.KeyByIdentifier(t.Identifier)
	a.Team = t
	a.TeamMap = p.Team
	a.UserMap = p.User
	a.Link = link(a.Identifier, p.WebHost)
	a.shortUser = p.ShortUser
	a.shortAlert = p.ShortAlert
	a.descriptionToName = p.DescriptionToName
	a.parseDescription = p.ParseDescription
	a.Name = a.findName()
	a.Prometheus = a.descriptionToDetail()

	return a
}
