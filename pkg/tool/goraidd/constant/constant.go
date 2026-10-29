package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goraidd",
	"RAID monitoring service",
	"goraidd [flags]",
)

const (
	HostEnvironment     = "GORAID_HOST"
	PortEnvironment     = "GORAID_PORT"
	InsecureEnvironment = "GORAID_INSECURE"
	TokenEnvironment    = "GORAID_TOKEN"
)

const (
	LogsTitle    = "Logs"
	LogsPath     = "/"
	ReportsTitle = "Reports"
	ReportsPath  = "/reports"
	RaidsTitle   = "Raids"
	RaidsPath    = "/raids"
	PlayersTitle = "Players"
	PlayersPath  = "/players"

	PageSize = 50

	EnrichCutoffDays     = 14
	FileDateLayout       = "20060102"
	FileDatePrefixLength = 8
)
