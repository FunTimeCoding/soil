package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gomonitord",
	"Shared claims for gomonitor windows",
	"gomonitord",
)

const (
	HostEnvironment     = "GOMONITOR_HOST"
	PortEnvironment     = "GOMONITOR_PORT"
	InsecureEnvironment = "GOMONITOR_INSECURE"
	TokenEnvironment    = "GOMONITOR_TOKEN" // #nosec G101 not a hardcoded secret
	ClaimsPath          = "/api/claims"
	StreamPath          = "/stream/claims"
	EventFormat         = "data: %s\n\n"
	EventPrefix         = "data: "
	ClaimRequiredText   = "item and owner are required"
)
