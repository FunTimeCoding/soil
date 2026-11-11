package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

const PollInterval = 5 * time.Minute

var Identity = identity.New(
	"gopatrold",
	"Agentless host health over SSH",
	"gopatrold [flags]",
)

const HostEnvironment = "PATROL_HOST"
const PatrolPeer = "patrol"
const (
	HostLabel  = "host"
	ArrayLabel = "array"
	DriveLabel = "drive"
	ModelLabel = "model"
)

const (
	ReachedMetric    = "patrol_host_reached"
	ReachedHelp      = "1 when the last round read the host"
	RequiredMetric   = "patrol_array_members_required"
	RequiredHelp     = "Members the array is built from"
	ActiveMetric     = "patrol_array_members_active"
	ActiveHelp       = "Members the array is running on"
	FailedMetric     = "patrol_array_members_failed"
	FailedHelp       = "Members the kernel marked failed"
	RebuildingMetric = "patrol_array_rebuilding"
	RebuildingHelp   = "1 while the array recovers, resyncs or reshapes"
)

const (
	PassedMetric      = "patrol_drive_health_passed"
	PassedHelp        = "1 when the drive's own health assessment passes"
	TemperatureMetric = "patrol_drive_temperature_celsius"
	TemperatureHelp   = "Drive temperature"
	WarningMetric     = "patrol_drive_critical_warning"
	WarningHelp       = "NVMe critical warning flags, 0 is none"
	SpareMetric       = "patrol_drive_available_spare_percent"
	SpareHelp         = "NVMe spare capacity left"
	UsedMetric        = "patrol_drive_used_percent"
	UsedHelp          = "NVMe endurance used, can pass 100"
	MediaMetric       = "patrol_drive_media_errors"
	MediaHelp         = "NVMe unrecovered data integrity errors"
)
