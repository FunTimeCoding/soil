package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

var Identity = identity.New(
	"goflightd",
	"macOS continuity flight recorder",
	"goflightd",
)

const (
	LogCommand      = "log"
	LogStream       = "stream"
	LogStyleFlag    = "--style"
	LogStyle        = "ndjson"
	LogInformation  = "--info"
	LogPredicate    = "--predicate"
	StreamPredicate = `(process == "ensembled" OR process == "rapportd" OR process == "sharingd" OR subsystem CONTAINS[c] "universalcontrol" OR subsystem CONTAINS[c] "ensemble") AND NOT (category IN {"BLEAdv", "WirelessProximity", "SDNearbyAgentCore"})`
	PredicateFlag   = "predicate"
	PredicateUsage  = "Log stream predicate"
	StreamTime      = "2006-01-02 15:04:05.999999-0700"

	Sudo             = "sudo"
	SudoNonInteract  = "-n"
	SystemProfiler   = "system_profiler"
	BluetoothSection = "SPBluetoothDataType"
	ProfilerNotation = "-json"

	WirelessKind  = "wireless"
	BluetoothKind = "bluetooth"
	KeySeparator  = "."

	Connected    = "connected"
	Disconnected = "disconnected"

	DateFormat = "2006-01-02T15:04:05Z07:00"

	LabelRequired = "label is required"

	StreamPause       = 5 * time.Second
	BluetoothInterval = 15 * time.Second
	WirelessInterval  = 5 * time.Second
	PruneInterval     = time.Hour
	Retention         = 7 * 24 * time.Hour
)
