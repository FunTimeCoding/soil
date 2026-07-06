package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goalert",
	"Prometheus alert check",
	"goalert [flags]",
)
