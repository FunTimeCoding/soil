package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gooutpostd",
	"Host service inventory with package provenance",
	"gooutpostd [flags]",
)

const (
	HostEnvironment     = "GOOUTPOST_HOST"
	PortEnvironment     = "GOOUTPOST_PORT"
	InsecureEnvironment = "GOOUTPOST_INSECURE"
	TokenEnvironment    = "GOOUTPOST_TOKEN"
)
