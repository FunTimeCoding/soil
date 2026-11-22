package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/types/measure"
	"time"
)

var Identity = identity.New(
	"gobatteryd",
	"Battery exporter over Network UPS Tools",
	"gobatteryd [flags]",
)

const (
	PollInterval = 30 * time.Second
	PeerFormat   = "battery server %s"
)

const (
	StatusVariable       = "ups.status"
	ModelVariable        = "device.model"
	ManufacturerVariable = "device.mfr"
	SerialVariable       = "device.serial"
)

const (
	StatusMetric      = "battery_status"
	StatusHelp        = "A battery's status flags, 1 for each flag it reports"
	InformationMetric = "battery_information"
	InformationHelp   = "A battery's model, manufacturer and serial, always 1"
	AttachedMetric    = "battery_attached"
	AttachedHelp      = "Whether the battery's driver serves its data, 1 or 0"
	DriverMetric      = "battery_driver_connected"
	DriverHelp        = "Whether the battery's driver is connected to upsd, stale or not, 1 or 0"
	ReachableMetric   = "battery_server_reachable"
	ReachableHelp     = "Whether an inventory server answered the last poll, 1 or 0"
)

const (
	ServerLabel       = "server"
	BatteryLabel      = "battery"
	FlagLabel         = "flag"
	ModelLabel        = "model"
	ManufacturerLabel = "manufacturer"
	SerialLabel       = "serial"
)

var Measures = []*measure.Measure{
	measure.New("battery.charge", "battery_charge_percent", "Battery charge"),
	measure.New(
		"battery.charge.low",
		"battery_charge_low_percent",
		"Charge at which the battery signals low battery",
	),
	measure.New(
		"battery.charge.warning",
		"battery_charge_warning_percent",
		"Charge at which the battery warns",
	),
	measure.New(
		"battery.runtime",
		"battery_runtime_seconds",
		"Runtime left on battery at the current load",
	),
	measure.New(
		"battery.runtime.low",
		"battery_runtime_low_seconds",
		"Runtime at which the battery signals low battery",
	),
	measure.New("ups.load", "battery_load_percent", "Load on the battery"),
	measure.New("input.voltage", "battery_input_volts", "Input voltage"),
	measure.New("output.voltage", "battery_output_volts", "Output voltage"),
}

var StatusWords = map[string]string{
	"OL":      "online",
	"OB":      "on_battery",
	"LB":      "low_battery",
	"HB":      "high_battery",
	"RB":      "replace_battery",
	"CHRG":    "charging",
	"DISCHRG": "discharging",
	"BYPASS":  "bypass",
	"CAL":     "calibrating",
	"OFF":     "offline",
	"OVER":    "overloaded",
	"TRIM":    "trimming",
	"BOOST":   "boosting",
	"FSD":     "forced_shutdown",
}
