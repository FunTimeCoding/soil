package constant

import "time"

const (
	PluginEnvironment = "MONITOR_PLUGINS"
	FileEnvironment   = "MONITOR_FILE"
	ManualEnvironment = "MONITOR_MANUAL"

	NotationReportLimit int = 10

	ReconnectDelay = 5 * time.Second
	OwnerFormat    = "%s@%s"
)
