package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gopackagedeb",
	"Debian package builder",
	"gopackagedeb <executable> <version>",
)

const (
	MaintainerNameEnvironment = "MAINTAINER_NAME"
	MaintainerMailEnvironment = "MAINTAINER_MAIL"
)
