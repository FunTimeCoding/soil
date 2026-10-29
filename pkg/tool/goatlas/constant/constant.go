package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goatlas",
	"Architecture documentation - what runs on which machine",
	"goatlas <command>",
)

const (
	PlaceArgument   = "place"
	PlaceUsage      = "Restrict to one device or virtual machine"
	SourceArgument  = "source"
	SourceUsage     = "Restrict to one attesting source"
	PackageArgument = "package"
	PackageUsage    = "Restrict to placements carrying this package"

	UnclaimedArgument = "unclaimed"
	UnclaimedUsage    = "Only hosts that resolve to no inventory object"
)
