package constant

import "time"

const (
	CruiseOff     = "off"
	CruiseDefault = "default"
	CruiseMinutes = "minutes"
)

const (
	CruiseInterval     = 30 * time.Second
	WakeWindow         = time.Hour
	DefaultCruiseGrace = 30 * time.Minute
)

const (
	CruiseTool = "cruise"
	Grace      = "grace"
	GraceKind  = "grace_kind"
)

const (
	CruiseGraceMinimumMinutes = 1
	CruiseGraceMaximumMinutes = 59
)

var CruiseKinds = []string{QueueMessage, QueueNotification, QueuePulse}
