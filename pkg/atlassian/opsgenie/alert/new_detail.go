package alert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert/option"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func NewDetail(
	v *alert.GetAlertResult,
	p *option.Alert,
	t *team.Team,
) *Alert {
	result := &Alert{
		Identifier:      v.Id,
		SmallIdentifier: strings.MustToInteger(v.TinyId),
		Status:          v.Status,
		Priority:        v.Priority,
		Seen:            v.IsSeen,
		Acknowledged:    v.Acknowledged,
		Create:          v.CreatedAt,
		Snoozed:         v.Snoozed,
		SnoozeUntil:     v.SnoozedUntil,
		Update:          v.UpdatedAt,
		Report:          v.Report,
		Responders:      v.Responders,
		Owner:           v.Owner,
		Source:          v.Source,
		Tags:            v.Tags,
		Details:         v.Details,
		Description:     v.Description,
		RawDetail:       v,
	}
	result.enrich(p, t)

	return result
}
