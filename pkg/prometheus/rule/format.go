package rule

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"slices"
	"time"
)

func (r *Rule) Format(f *option.Format) string {
	s := status.New(f).String(
		r.formatName(f),
		r.Group,
		r.formatType(),
	)

	if r.Summary != "" {
		s.Line("  Summary: %s", r.Summary)
	}

	if r.Description != "" {
		s.Line("  Description: %s", r.Description)
	}

	if r.Duration > 0 {
		s.Line("  Duration: %d", r.Duration)
	}

	if !slices.Contains(Healths, r.Health) {
		output := fmt.Sprintf("unexpected health: %s", r.Health)

		if f.UseColor {
			output = console.Red("%s", output)
		}

		s.Line("%s", output)
	}

	if !slices.Contains(States, r.State) {
		output := fmt.Sprintf("unexpected state: %s", r.State)

		if f.UseColor {
			output = console.Red("%s", output)
		}

		s.Line("%s", output)
	}

	if r.RawAlert != nil {
		if age := time.Since(
			r.RawAlert.LastEvaluation,
		).Round(time.Second); age > 30*time.Second {
			output := fmt.Sprintf("  Alert last evaluation: %s", age)

			if f.UseColor {
				output = console.Red("%s", output)
			}

			s.Line("%s", output)
		}

		if r.RawAlert.EvaluationTime > 0.1 {
			output := fmt.Sprintf(
				"  EvaluationTime: %.1f",
				r.RawAlert.EvaluationTime,
			)

			if f.UseColor {
				output = console.Yellow("%s", output)
			}

			s.Line("%s", output)
		}

		if false {
			if false {
				r.RawAlert.LastEvaluation = time.Time{}
			}

			r.RawAlert.Name = ""
			r.RawAlert.State = ""
			r.RawAlert.Health = ""
			r.RawAlert.Duration = 0
			r.RawAlert.Query = ""
			delete(r.RawAlert.Labels, SeverityKey)
			delete(r.RawAlert.Annotations, SummaryKey)
			delete(r.RawAlert.Annotations, DescriptionKey)
			delete(r.RawAlert.Annotations, RunbookKey)
			delete(r.RawAlert.Annotations, DurationKey)
		}
	}

	s.RawList(r)

	if r.RawRecord != nil {
		s.Raw(r.RawRecord, "RawRecord")
	}

	if r.RawAlert != nil {
		s.Raw(r.RawAlert, "RawAlert")
	}

	if r.RawGroup != nil {
		s.Raw(r.RawGroup, "RawGroup")
	}

	return s.Format()
}
