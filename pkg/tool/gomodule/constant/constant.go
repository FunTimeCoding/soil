package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gomodule",
	"Go module maintenance: dependency update, version, license and staleness checks",
	"gomodule <command> [flags]",
)

const (
	UpdateUsage  = "update"
	VersionUsage = "version [path]"
	LicenseUsage = "license [path]"
	StaleUsage   = "stale [path]"
)

const (
	DefaultDepth  = 3
	DefaultMonths = 12
)
