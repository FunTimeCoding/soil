package rule_list

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

func alertDrifted(
	a *v1.AlertingRule,
	b *v1.AlertingRule,
) bool {
	if !LabelsSame(a.Labels, b.Labels) || a.Query != b.Query {
		return false
	}

	return !cmp.Equal(
		a,
		b,
		cmpopts.IgnoreFields(
			v1.AlertingRule{},
			"EvaluationTime",
			"LastEvaluation",
		),
	)
}
