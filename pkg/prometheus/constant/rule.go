package constant

import "github.com/prometheus/client_golang/api/prometheus/v1"

const (
	SeverityKey = "severity"

	SummaryKey       = "summary"
	DescriptionKey   = "description"
	DurationKey      = "duration"
	RunbookKey       = "runbook_url"
	DocumentationKey = "docs"

	AlertType                 = "alert"
	RecordType                = "record"
	AllType                   = "all"
	UnknownType               = "unknown"
	HealthOkay  v1.RuleHealth = "ok"

	InactiveState = "inactive"
	PendingState  = "pending"
	FiringState   = "firing"
)

var (
	RuleHealths = []v1.RuleHealth{
		HealthOkay,
	}
	RuleStates = []string{
		InactiveState,
		PendingState,
		FiringState,
	}
)
