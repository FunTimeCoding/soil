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

	GeneratePath   = "/generate"
	CreateRaidPath = "/raids/create"
	StaticPath     = "/static/"
	DeleteSuffix   = "/delete"

	FileNameField       = "fileNames"
	FileNameParameter   = "fileName"
	IdentifierParameter = "id"
	AccountParameter    = "account"
	AccountColumn       = "account"
	OffsetParameter     = "offset"
	StartParameter      = "start"
	EndParameter        = "end"
	FilterFormat        = "2006-01-02T15:04"

	ReportSizeClass = "report-size"

	PageSize = 50

	EnrichCutoffDays     = 14
	FileDateLayout       = "20060102"
	FileDatePrefixLength = 8
)
