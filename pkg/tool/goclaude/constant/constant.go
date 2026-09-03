package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"regexp"
)

var Identity = identity.New(
	"goclaude",
	"Claude Code session management and analysis",
	"goclaude [command]",
)

const (
	NameEnvironment  = "CLAUDE_NAME"
	HostEnvironment  = "CLAUDE_HOST"
	PortEnvironment  = "CLAUDE_PORT"
	TokenEnvironment = "CLAUDE_TOKEN"
	PeekOutputBudget = 120
	PeekContextLimit = 200

	EnvironmentFileEnvironment   = "CLAUDE_ENV_FILE"
	SessionIdentifierEnvironment = "CLAUDE_SESSION_ID"

	GuardBlockExit = 2
	SedMessage     = "sed on macOS is BSD sed and its flags (notably -i) differ from GNU sed - use gsed instead"
	NpxMessage     = "npx is blocked (supply-chain guard) - it downloads and executes npm packages on demand"
	PipMessage     = "pip install is blocked (supply-chain guard) - no python dependencies may be installed on this system"

	NoGuardEnvironment = "CLAUDE_NO_GUARD"

	StatusLineDumpFile = "/tmp/goclaude-status-line.json"
)

var (
	SedInvocation = regexp.MustCompile(`(^|[|&;(\s])sed(\s|$)`)
	NpxInvocation = regexp.MustCompile(`(^|[|&;(\s])npx(\s|$)`)
	PipInvocation = regexp.MustCompile(
		`(^|[|&;(\s])pip3?\s(.*\s)?install(\s|$)`,
	)
	ShortModelName = map[string]string{"Fable 5": "Fable"}
)
