package constant

import "time"

const (
	Port    = 3493
	Timeout = 5 * time.Second
	Network = "tcp"
)

const (
	ListBatteries  = "LIST UPS"
	ListVariables  = "LIST VAR %s"
	Logout         = "LOGOUT"
	ExchangeFormat = "%s\n%s\n"
)

const (
	BatteryLine    = "UPS"
	VariableLine   = "VAR"
	EndPrefix      = "END LIST"
	ErrorPrefix    = "ERR "
	UnknownBattery = "UNKNOWN-UPS"
	RefusedFormat  = "%s refused: %s"
	BatteryKind    = "battery"
)

const (
	StaleReadings      = "DATA-STALE"
	DriverNotConnected = "DRIVER-NOT-CONNECTED"
	DetachedFormat     = "battery %s detached: %s"
)
