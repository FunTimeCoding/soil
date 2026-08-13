package alert

import (
	"github.com/docker/go-units"
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/constant"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings/join"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"time"
)

func (a *Alert) extended(
	s *status.Status,
	f *option.Format,
) {
	s.Line("  %s", a.Link)

	if f.HasTag(consoleConstant.TagName) {
		s.Line("  Name: %s", a.Name)
	}

	investigate := f.HasTag(consoleConstant.TagInvestigate)

	if investigate && a.Update.Sub(a.Create) > time.Minute {
		s.Line("  Update: %s", condenseTime(a.Update))
	}

	if a.Snoozed {
		s.Line("  Snoozed: %s", a.SnoozeUntil.Format(timeConstant.DateMinute))
	}

	if investigate && a.Report.AckTime > 0 {
		s.Line(
			"  AckTime: %s",
			units.HumanDuration(
				time.Duration(a.Report.AckTime/1000)*time.Second,
			),
		)
	}

	if a.Report.AcknowledgedBy != a.Owner {
		by := a.shortenUser(a.Report.AcknowledgedBy)

		if f.UseColor {
			by = consoleConstant.Yellow("%s", by)
		}

		s.Line("  Acknowledged: %s", by)
	}

	if a.Report.CloseTime > 0 {
		s.Line("  Closed: %s", a.Report.ClosedBy)
	}

	dense := f.HasTag(consoleConstant.TagDense)

	if !dense {
		var responders []string

		for _, r := range a.Responders {
			if r.Type == alert.TeamResponder {
				var name string

				if t := a.TeamMap.ByIdentifier(r.Id); t != nil {
					name = a.TeamMap.KeyByName(t.Name)

					if name == atlassian.OpsgenieNoKey {
						name = atlassian.OpsgenieUnknownTeam
					}
				} else {
					name = atlassian.OpsgenieUnknownTeam
				}

				if name == atlassian.OpsgenieUnknownTeam {
					s.Line("  Unknown responder team: %+v", r)
				} else {
					responders = append(responders, name)
				}
			} else {
				var name string

				if u := a.UserMap.ByIdentifier(r.Id); u != nil {
					name = a.shortenUser(u.Name)

					if name == atlassian.OpsgenieNoKey {
						name = atlassian.OpsgenieUnknownUser
					}
				} else {
					name = atlassian.OpsgenieUnknownUser
				}

				if name == atlassian.OpsgenieUnknownUser {
					s.Line("  Unknown responder user: %+v", r)
				} else {
					responders = append(responders, name)
				}
			}
		}

		if len(responders) > 0 {
			s.Line("  Responders: %s", join.Comma(responders))
		}

		if len(a.Tags) > 0 {
			s.Line("  Tags: %s", join.Comma(a.Tags))
		}

		if a.Source != "" {
			s.Line("  Source: %s", a.Source)
		}
	}
}
