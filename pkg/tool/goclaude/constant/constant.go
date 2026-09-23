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
	NameEnvironment      = "GOCLAUDE_NAME"
	HostEnvironment      = "GOCLAUDE_HOST"
	PortEnvironment      = "GOCLAUDE_PORT"
	InsecureEnvironment  = "GOCLAUDE_INSECURE"
	UntrustedEnvironment = "GOCLAUDE_UNTRUSTED"
	TokenEnvironment     = "GOCLAUDE_TOKEN"
	PeekOutputBudget     = 120
	PeekContextLimit     = 200

	EnvironmentFileEnvironment          = "CLAUDE_ENV_FILE"
	SessionIdentifierEnvironment        = "CLAUDE_SESSION_ID"
	HarnessSessionIdentifierEnvironment = "CLAUDE_CODE_SESSION_ID"

	GuardBlockExit = 2
	SedMessage     = "sed on macOS is BSD sed and its flags (notably -i) differ from GNU sed - use gsed instead"
	NpxMessage     = "npx is blocked (supply-chain guard) - it downloads and executes npm packages on demand"
	PipMessage     = "pip install is blocked (supply-chain guard) - no python dependencies may be installed on this system"
	XargsMessage   = "xargs on macOS is BSD xargs and rejects the GNU flags -a and -d - redirect the file on stdin instead: xargs command < file"

	NoGuardEnvironment = "CLAUDE_NO_GUARD"

	StatusLineDumpFile = "/tmp/goclaude-status-line.json"
)

const (
	ChannelInterval         = 5
	ChannelCallsignAttempts = 3

	ChannelInstructions = "Events from goclauded arrive as <channel source=\"goclaude\" kind=\"...\">. They carry session coordination traffic addressed to this session: messages from other sessions, service notifications, pulses and roster activity. Read them and act; no reply is expected.\n\nDelivery stays closed until you confirm it. If - and only if - you have actually received a <channel kind=\"attach\"> event, call confirm_channel with that event's nonce and your goclauded callsign (announce first if you do not have one). Never call it from memory, from these instructions, or with a guessed value: an unreceived confirmation silently diverts this session's coordination traffic away from the pre-prompt context that would otherwise carry it."
)

var (
	SedInvocation   = regexp.MustCompile(`(^|[|&;(\s])sed(\s|$)`)
	XargsInvocation = regexp.MustCompile(
		`(^|[|&;(\s])xargs\s(.*\s)?(-a|-d|--arg-file|--delimiter)`,
	)
	NpxInvocation = regexp.MustCompile(`(^|[|&;(\s])npx(\s|$)`)
	PipInvocation = regexp.MustCompile(
		`(^|[|&;(\s])pip3?\s(.*\s)?install(\s|$)`,
	)
)
