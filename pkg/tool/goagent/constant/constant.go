package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goagent",
	"CLI for the goagentd agent runner",
	"goagent [command]",
)
