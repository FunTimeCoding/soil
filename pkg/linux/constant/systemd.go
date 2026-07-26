package constant

const (
	SystemdCommand   = "systemctl"
	SystemdListUnits = "list-units"
	SystemdStatus    = "status"
	SystemdShow      = "show"

	SystemdNoLegend = "--no-legend"

	SystemdAll   = "--all"   // Units in memory, including dead and empty
	SystemdFull  = "--full"  // Do not shorten unit names
	SystemdPlain = "--plain" // Dependencies as a list instead of tree

	SystemdState    = "--state"
	SystemdNotFound = "not-found"

	SystemdType    = "--type"
	SystemdService = "service"

	SystemdOutput   = "--output"
	SystemdNotation = "json"

	SystemdDateTime               = "Mon 2006-01-02 15:04:05 MST"
	SystemdActiveState            = "active"
	SystemdFailedState            = "failed"
	SystemdRunningSubState        = "running"
	SystemdFailedSubState         = "failed"
	SystemdActiveEnterTimestamp   = "ActiveEnterTimestamp"
	SystemdExecMainStartTimestamp = "ExecMainStartTimestamp"
)
