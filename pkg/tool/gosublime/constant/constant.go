package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gosublime",
	"Sublime Text buffers - list, read, create, edit, open, save, close",
	"gosublime <command>",
)

const (
	HostEnvironment = "GOSUBLIMED_HOST"
	DefaultHost     = "localhost:8580"
)
