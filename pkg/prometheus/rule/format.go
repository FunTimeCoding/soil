package rule

import (
	"fmt"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"slices"
	"time"
)

func (r *Rule) Format(f *option.Format) string {
	s := status.New(f).String(r.formatName(f), r.Group, r.formatType())

	if r.Summary != "" {
		s.Line("  Summary: %s", r.Summary)
	}

	if r.Description != "" {
		s.Line("  Description: %s", r.Description)
	}

	if r.Duration > 0 {
		s.Line("  Duration: %d", r.Duration)
	}

	if !slices.Contains(constant.RuleHealths, r.Health) {
		output := fmt.Sprintf("unexpected health: %s", r.Health)

		if f.UseColor {
			output = consoleConstant.Red("%s", output)
		}

		s.Line("%s", output)
	}

	if !slices.Contains(constant.RuleStates, r.State) {
		output := fmt.Sprintf("unexpected state: %s", r.State)

		if f.UseColor {
			output = consoleConstant.Red("%s", output)
		}

		s.Line("%s", output)
	}

	if r.RawAlert != nil {
		if age := time.Since(r.RawAlert.LastEvaluation).Round(time.Second); age > 30*time.Second {
			output := fmt.Sprintf("  Alert last evaluation: %s", age)

			if f.UseColor {
				output = consoleConstant.Red("%s", output)
			}

			s.Line("%s", output)
		}

		if r.RawAlert.EvaluationTime > 0.1 {
			output := fmt.Sprintf(
				"  EvaluationTime: %.1f",
				r.RawAlert.EvaluationTime,
			)

			if f.UseColor {
				output = consoleConstant.Yellow("%s", output)
			}

			s.Line("%s", output)
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
