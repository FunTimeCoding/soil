package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

var Identity = identity.New("godashboardd", "Homelab dashboard", "godashboardd")

const (
	DashboardTitle = "Dashboard"
	DashboardPath  = "/"
	HeatmapTitle   = "Heatmap"
	HeatmapPath    = "/heatmap"

	BoardUsage = "Board definition file path"
	BoardFile  = "godashboardd.yaml"

	PendingValue = "–"

	RefreshInterval = 30 * time.Second
)
const (
	IconHost          = "https://cdn.jsdelivr.net/gh/homarr-labs/dashboard-icons/png/"
	FilesLabel        = "Files"
	SharesLabel       = "Shares"
	ApplicationsLabel = "Apps"
	OutOfSyncLabel    = "OutOfSync"
	DegradedLabel     = "Degraded"
	MissingLabel      = "Missing"
	RowEventPrefix    = "rows-"
	LabelAttribute    = "data-label"
)
const LabelColumn = "label"
const (
	DefaultTailColumns = 4
	NextcloudWidget    = "nextcloud"
	ArgocdWidget       = "argocd"
	SecurePort         = 443
	InsecurePort       = 80
)
