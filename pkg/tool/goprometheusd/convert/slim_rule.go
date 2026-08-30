package convert

type SlimRule struct {
	Name           string          `json:"name"`
	Group          string          `json:"group"`
	Type           string          `json:"type"`
	State          string          `json:"state,omitempty"`
	Health         string          `json:"health"`
	Query          string          `json:"query"`
	Severity       string          `json:"severity,omitempty"`
	Duration       int             `json:"duration,omitzero"`
	LastEvaluation string          `json:"last_evaluation,omitempty"`
	EvaluationTime float64         `json:"evaluation_time,omitzero"`
	LastError      string          `json:"last_error,omitempty"`
	Alerts         []SlimRuleAlert `json:"alerts,omitempty"`
}
