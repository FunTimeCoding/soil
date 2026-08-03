package constant

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	constant1 "github.com/funtimecoding/soil/pkg/prometheus/constant"
)

const (
	OpsgenieUserKeyEnvironment = "OPSGENIE_USER_KEY"
	OpsgenieTeamKeyEnvironment = "OPSGENIE_TEAM_KEY"
	OpsgenieWebHostEnvironment = "OPSGENIE_WEB_HOST"
	OpsgenieTeamEnvironment    = "OPSGENIE_TEAM"

	OpsgeniePageLimit int = 100

	OpsgenieNoKey = "no key"
)

var OpsgenieFormat = constant.ExtendedColorFormat.Copy().Tag(
	constant.TagCategory,
)

const (
	OpsgenieNoOwner = "no owner"
	OpsgenieNoTeam  = "no team"

	OpsgenieUnknownName = "unknown name"
	OpsgenieUnknownTeam = "unknown team"
	OpsgenieUnknownUser = "unknown user"

	OpsgenieClosedStatus = "closed"
	OpsgenieOpenStatus   = "open"

	OpsgeniePriorityP1 = "P1"
	OpsgeniePriorityP2 = "P2"
	OpsgeniePriorityP3 = "P3"
	OpsgeniePriorityP4 = "P4"
	OpsgeniePriorityP5 = "P5"

	OpsgenieUnacknowledgedFlag = "unacknowledged"
	OpsgenieSnoozedFlag        = "snoozed"
	OpsgenieUnseenFlag         = "unseen"
)

var (
	OpsgenieStatuses = []string{
		OpsgenieClosedStatus,
		OpsgenieOpenStatus,
	}
	OpsgeniePriorities = []string{
		OpsgeniePriorityP1,
		OpsgeniePriorityP2,
		OpsgeniePriorityP3,
		OpsgeniePriorityP4,
		OpsgeniePriorityP5,
	}
	OpsgenieSkipDetail = []string{
		constant1.SeverityLabel,
	}
	OpsgenieCondenseSkipFields []string
)
