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

	IconHost          = "https://cdn.jsdelivr.net/gh/homarr-labs/dashboard-icons/png/"
	FilesLabel        = "Files"
	SharesLabel       = "Shares"
	ApplicationsLabel = "Apps"
	OutOfSyncLabel    = "OutOfSync"
	DegradedLabel     = "Degraded"
	MissingLabel      = "Missing"
	RowEventPrefix    = "rows-"
	LabelAttribute    = "data-label"

	BoardColumnsVariable = "--board-columns"

	LabelColumn = "label"

	DefaultTailColumns = 4
	NextcloudWidget    = "nextcloud"
	ArgocdWidget       = "argocd"
	SecurePort         = 443
	InsecurePort       = 80

	PrometheusPeer = "prometheus"
	ArgocdPeer     = "argocd"
	NextcloudPeer  = "nextcloud"
	RefreshFailed  = "refresh deferred: %v"
)

const ArgocdTokenEnvironment = "ARGOCD_TOKEN"
