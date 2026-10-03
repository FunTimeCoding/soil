package constant

const (
	SocketEnvironment = "SSH_AUTH_SOCK"

	Command                  = "ssh"
	LocalPortForwardArgument = "-L"
	NoRemoteCommandArgument  = "-N"
	BackgroundArgument       = "-f"
	NoPTYArgument            = "-T"
	ForcePTYArgument         = "-tt"
	VerboseArgument          = "-v"

	TerminalType     = "xterm"
	TerminalHeight   = 25
	TerminalWidth    = 80
	TerminalBaudRate = 14400
)
const (
	TargetHost = "target-host"
	TargetPort = "target-port"
)
