package alert

import (
	"fmt"
	"github.com/docker/go-units"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/time"
)

func (a *Alert) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Fingerprint) {
		s.String(a.Fingerprint)
	}

	s.String(a.formatEntity(f))

	if f.HasTag(tag.Category) {
		s.String(a.formatCategory(f))
	}

	linkLabel := a.Name

	if a.Entity == "" && a.Category == "" {
		linkLabel = "Alert"
		s.String(a.formatName(f))
	}

	s.String(a.formatSeverity(f))

	if f.HasTag(tag.Host) {
		s.String(a.formatHost(f))
	}

	if f.HasTag(tag.Instance) {
		s.String(a.formatInstance())
	}

	if a.Start != nil {
		if false {
			s.String(a.Start.Format(time.DateMinute))
		}

		s.String(fmt.Sprintf("%s ago", units.HumanDuration(a.Age())))
	}

	s.String(a.formatConcern(f))
	s.DetailLink(a.Link, linkLabel, "")

	if a.Runbook != constant.None {
		s.DetailLink(a.Runbook, "Runbook", "Runbook")
	}

	for _, e := range a.ExtraBubble {
		s.String(e)
	}

	if f.ShowExtended {
		if a.Summary != constant.None {
			s.Line("  Summary: %s", a.Summary)
		}

		if a.Message != constant.None {
			s.Line("  Message: %s", a.Message)
		}

		if v := a.formatRemainingLabels(f); v != "" {
			s.Line("  Labels: %s", v)
		}

		if len(a.Receivers) > 0 {
			s.Line("  Receivers: %s", join.Comma(a.Receivers))
		}

		if a.HostLink != "" {
			if f.HasTag(tag.Copyable) || f.HasTag(tag.Markdown) {
				s.DetailLink(a.HostLink, "Host", "Host")
			}
		}
	}

	s.RawList(a.Raw)

	return s.Format()
}
