package alert

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"

const (
	NoOwner = "no owner"
	NoTeam  = "no team"

	UnknownName = "unknown name"
	UnknownTeam = "unknown team"
	UnknownUser = "unknown user"

	ClosedStatus = "closed"
	OpenStatus   = "open"

	PriorityP1 = "P1"
	PriorityP2 = "P2"
	PriorityP3 = "P3"
	PriorityP4 = "P4"
	PriorityP5 = "P5"

	UnacknowledgedFlag = "unacknowledged"
	SnoozedFlag        = "snoozed"
	UnseenFlag         = "unseen"
)

var (
	Statuses = []string{
		ClosedStatus,
		OpenStatus,
	}
	Priorities = []string{
		PriorityP1,
		PriorityP2,
		PriorityP3,
		PriorityP4,
		PriorityP5,
	}
	SkipDetail = []string{
		constant.SeverityLabel,
	}
	CondenseSkipFields []string
)
