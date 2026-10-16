package silence

import (
	"fmt"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/openapi"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func New(
	v *models.GettableSilence,
	link *locator.Locator,
) *Silence {
	var match []string
	var rule string

	for _, m := range v.Matchers {
		if *m.Name == constant.AlertnameLabel {
			rule = *m.Value
		}

		match = append(match, fmt.Sprintf("%s=%s", *m.Name, *m.Value))
	}

	if rule == "" {
		rule = constant.UnknownRule
	}

	return &Silence{
		MonitorIdentifier: monitorConstant.GoSilence.StringIdentifier(*v.ID),
		Identifier:        *v.ID,
		State:             *v.Status.State,
		Match:             join.Comma(match),
		Start:             openapi.ConvertTime(v.StartsAt),
		End:               openapi.ConvertTime(v.EndsAt),
		Author:            *v.CreatedBy,
		Comment:           *v.Comment,
		Rule:              rule,
		Link: link.Copy().Trail().Fragment(
			"/silences/%s",
			*v.ID,
		).String(),
		Raw: v,
	}
}
