package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goagentd",
	"Agent runner daemon - receives intents and executes Claude Code",
	"goagentd",
)

const (
	HostEnvironment     = "GOAGENT_HOST"
	PortEnvironment     = "GOAGENT_PORT"
	InsecureEnvironment = "GOAGENT_INSECURE"
	TokenEnvironment    = "GOAGENT_TOKEN"
)
