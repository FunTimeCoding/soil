package constant

import "github.com/gorilla/websocket"

const (
	PluginEnvironment = "MONITOR_PLUGINS"
	FileEnvironment   = "MONITOR_FILE"
	ManualEnvironment = "MONITOR_MANUAL"

	NotationReport int = 10 // Limit
)

// Command
const (
	LoginCommand  = "login"
	LogoutCommand = "logout"
	FlagCommand   = "flag"
	ClearCommand  = "clear"
	PingCommand   = "ping"

	LoginResponseCommand = "login-response"
	FlagAddCommand       = "flag-add"
	FlagRemoveCommand    = "flag-remove"
)

var (
	Upgrader = websocket.Upgrader{}
)
