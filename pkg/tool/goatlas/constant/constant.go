package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goatlas",
	"Architecture documentation - what runs on which machine",
	"goatlas <command>",
)

const (
	HostEnvironment     = "GOATLAS_HOST"
	PortEnvironment     = "GOATLAS_PORT"
	InsecureEnvironment = "GOATLAS_INSECURE"
	TokenEnvironment    = "GOATLAS_TOKEN"

	PlaceArgument  = "place"
	PlaceUsage     = "Restrict to one device or virtual machine"
	SourceArgument = "source"
	SourceUsage    = "Restrict to one attesting source"

	Unplaced  = "-"
	Unclaimed = "unclaimed"

	UnclaimedArgument = "unclaimed"
	UnclaimedUsage    = "Only hosts that resolve to no inventory object"
	ScopeSeparator    = "/"
)
