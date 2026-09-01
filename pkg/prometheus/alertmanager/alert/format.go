package alert

import (
	"fmt"
	"github.com/docker/go-units"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
)

func (a *Alert) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(consoleConstant.TagFingerprint) {
		s.String(a.Fingerprint)
	}

	s.String(a.formatEntity(f))

	if f.HasTag(consoleConstant.TagCategory) {
		s.String(a.formatCategory(f))
	}

	linkLabel := a.Name

	if a.Entity == "" && a.Category == "" {
		linkLabel = "Alert"
		s.String(a.formatName(f))
	}

	s.String(a.formatSeverity(f))

	if f.HasTag(consoleConstant.TagHost) {
		s.String(a.formatHost(f))
	}

	if f.HasTag(consoleConstant.TagInstance) {
		s.String(a.formatInstance())
	}

	if a.Start != nil {
		if false {
			s.String(a.Start.Format(timeConstant.DateMinute))
		}

		s.String(fmt.Sprintf("%s ago", units.HumanDuration(a.Age())))
	}

	s.String(a.formatConcern(f))
	s.DetailLink(a.Link, linkLabel, "")

	if a.Runbook != prometheus.None {
		s.DetailLink(a.Runbook, "Runbook", "Runbook")
	}

	for _, e := range a.ExtraBubble {
		s.String(e)
	}

	if f.ShowExtended {
		if a.Summary != prometheus.None {
			s.Line("  Summary: %s", a.Summary)
		}

		if a.Message != prometheus.None {
			s.Line("  Message: %s", a.Message)
		}

		if v := a.formatRemainingLabels(f); v != "" {
			s.Line("  Labels: %s", v)
		}

		if len(a.Receivers) > 0 {
			s.Line("  Receivers: %s", join.Comma(a.Receivers))
		}

		if a.HostLink != "" {
			if f.HasTag(consoleConstant.TagCopyable) || f.HasTag(consoleConstant.TagMarkdown) {
				s.DetailLink(a.HostLink, "Host", "Host")
			}
		}
	}

	s.RawList(a.Raw)

	return s.Format()
}
