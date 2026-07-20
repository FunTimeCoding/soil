package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"godebian",
	"Debian package builder",
	"godebian <executable> <version>",
)

const (
	MaintainerNameArgument  = "maintainer-name"
	MaintainerEmailArgument = "maintainer-email"
	SystemdUnitFlag         = "systemd-unit"
)
