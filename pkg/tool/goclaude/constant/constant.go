package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goclaude",
	"Claude Code session management and analysis",
	"goclaude [command]",
)

const (
	NameEnvironment  = "CLAUDE_NAME"
	HostEnvironment  = "CLAUDE_HOST"
	PortEnvironment  = "CLAUDE_PORT"
	PeekOutputBudget = 120
	PeekContextLimit = 200

	EnvironmentFileEnvironment   = "CLAUDE_ENV_FILE"
	SessionIdentifierEnvironment = "CLAUDE_SESSION_ID"
)
