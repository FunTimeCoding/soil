package rule_list

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

func recordDrifted(
	a *v1.RecordingRule,
	b *v1.RecordingRule,
) bool {
	if !LabelsSame(a.Labels, b.Labels) || a.Query != b.Query {
		return false
	}

	return !cmp.Equal(
		a,
		b,
		cmpopts.IgnoreFields(
			v1.RecordingRule{},
			"EvaluationTime",
			"LastEvaluation",
		),
	)
}
