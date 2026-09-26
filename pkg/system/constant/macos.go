package constant

const (
	Codesign = "codesign"

	Launchctl      = "launchctl"
	Bootout        = "bootout"
	Bootstrap      = "bootstrap"
	LaunchctlPrint = "print"

	Wdutil            = "wdutil"
	WdutilInformation = "info"
)

const NotAvailable = "n/a"
const (
	LaunchctlList    = "list"
	LaunchctlHeader  = "PID"
	LaunchctlAbsent  = "-"
	LaunchctlSuccess = "0"

	LaunchAgentDirectory  = "/Library/LaunchAgents"
	LaunchDaemonDirectory = "/Library/LaunchDaemons"
	UserLaunchAgents      = "Library/LaunchAgents"
	SystemLaunchPrefix    = "/System/"
	PropertyListExtension = ".plist"
)
