package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goalertlogd",
	"Alert lifecycle tracker",
	"goalertlogd",
).WithInstructions(
	"Alert lifecycle tracker - query active, recent, and top alerts. Tracks when alerts start, end, and recur.",
)

const (
	GetAlerts       = "get_alerts"
	GetRecentAlerts = "get_recent_alerts"
	GetTopAlerts    = "get_top_alerts"
	GetStatus       = "get_status"

	End   = "end"
	N     = "n"
	Start = "start"

	DashboardTitle = "Dashboard"
	DashboardPath  = "/"
	RecentTitle    = "Recent"
	RecentPath     = "/recent"
)
