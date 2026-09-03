package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gosublime",
	"Sublime Text buffers - list, read, create, edit, open, save, close",
	"gosublime <command>",
)

const (
	HostEnvironment = "GOSUBLIMED_HOST"
	TokenEnvironment = "GOSUBLIMED_TOKEN"
	DefaultHost     = "localhost:8580"
)
