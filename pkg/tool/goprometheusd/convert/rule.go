package convert

import (
	"github.com/funtimecoding/soil/pkg/prometheus/rule"
	"time"
)

func Rule(r *rule.Rule) *SlimRule {
	result := &SlimRule{
		Name:     r.Name,
		Group:    r.Group,
		State:    r.State,
		Health:   string(r.Health),
		Query:    r.Query,
		Severity: r.Severity,
		Duration: r.Duration,
	}

	if r.RawAlert != nil {
		result.Type = rule.AlertType
		result.LastEvaluation = r.RawAlert.LastEvaluation.Format(time.RFC3339)
		result.EvaluationTime = r.RawAlert.EvaluationTime
		result.LastError = r.RawAlert.LastError

		for _, a := range r.RawAlert.Alerts {
			result.Alerts = append(
				result.Alerts,
				SlimRuleAlert{
					State:       string(a.State),
					ActiveAt:    a.ActiveAt.Format(time.RFC3339),
					Labels:      labelSetToMap(a.Labels),
					Annotations: labelSetToMap(a.Annotations),
					Value:       a.Value,
				},
			)
		}
	} else if r.RawRecord != nil {
		result.Type = rule.RecordType
		result.LastEvaluation = r.RawRecord.LastEvaluation.Format(time.RFC3339)
		result.EvaluationTime = r.RawRecord.EvaluationTime
		result.LastError = r.RawRecord.LastError
	}

	return result
}
