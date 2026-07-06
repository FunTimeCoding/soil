package option

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team_map"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/user_map"
	"github.com/funtimecoding/soil/pkg/face"
)

func New(
	t *team_map.Map,
	u *user_map.Map,
	webHost string,
	alert face.StringAlias,
	user face.StringAlias,
	descriptionToName face.StringAlias,
	parseDescription constant.ParseDescription,
) *Alert {
	return &Alert{
		Team:              t,
		User:              u,
		WebHost:           webHost,
		ShortAlert:        alert,
		ShortUser:         user,
		DescriptionToName: descriptionToName,
		ParseDescription:  parseDescription,
	}
}
