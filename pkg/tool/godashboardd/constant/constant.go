package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

var Identity = identity.New(
	"godashboardd",
	"Homelab dashboard",
	"godashboardd",
)

const (
	DashboardTitle = "Dashboard"
	DashboardPath  = "/"
	HeatmapTitle   = "Heatmap"
	HeatmapPath    = "/heatmap"

	BoardUsage = "Board definition file path"
	BoardFile  = "godashboardd.yaml"

	SignInPath   = "/sign-in"
	CallbackPath = "/callback"
	SignOutPath  = "/sign-out"

	IssuerEnvironment           = "AUTHORIZATION_ISSUER"
	ClientIdentifierEnvironment = "AUTHORIZATION_CLIENT_IDENTIFIER"
	ClientSecretEnvironment     = "AUTHORIZATION_CLIENT_SECRET"
	EncryptionSecretEnvironment = "AUTHORIZATION_ENCRYPTION_SECRET"
	PublicLocatorEnvironment    = "PUBLIC_LOCATOR"

	PendingValue = "–"

	RefreshInterval = 30 * time.Second
)
const (
	PalettePath       = "/palette"
	NavigateCategory  = "navigate"
	LinkCategory      = "link"
	IconHost          = "https://cdn.jsdelivr.net/gh/homarr-labs/dashboard-icons/png/"
	FilesLabel        = "Files"
	SharesLabel       = "Shares"
	ApplicationsLabel = "Apps"
	OutOfSyncLabel    = "OutOfSync"
	DegradedLabel     = "Degraded"
	MissingLabel      = "Missing"
	EventPath         = "/event"
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
