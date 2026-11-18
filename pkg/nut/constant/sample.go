package constant

const (
	FixtureGoodbye        = "OK Goodbye\n"
	FixtureUnknownCommand = "ERR UNKNOWN-COMMAND\n"
)

const FixtureBatteries = `BEGIN LIST UPS
UPS alfa "Description unavailable"
UPS bravo "Description unavailable"
END LIST UPS
`

const FixtureAlfaVariables = `BEGIN LIST VAR alfa
VAR alfa battery.charge "100"
VAR alfa battery.charge.low "10"
VAR alfa battery.runtime "3061"
VAR alfa battery.runtime.low "120"
VAR alfa ups.load "16"
VAR alfa input.voltage "239.0"
VAR alfa ups.status "OL CHRG"
VAR alfa device.model "Charlie \"Delta\" 850 "
VAR alfa device.mfr "Echo"
VAR alfa device.serial "9B00"
END LIST VAR alfa
`

const FixtureAlfaBattery = `BEGIN LIST UPS
UPS alfa "Description unavailable"
END LIST UPS
`

const FixtureBravoBattery = `BEGIN LIST UPS
UPS bravo "Description unavailable"
END LIST UPS
`

const FixtureBravoVariables = `BEGIN LIST VAR bravo
VAR bravo battery.charge "40"
VAR bravo ups.status "OB LB FOXTROT"
END LIST VAR bravo
`

const (
	FixtureUnknownBattery = "ERR UNKNOWN-UPS\n"
	FixtureAccessDenied   = "ERR ACCESS-DENIED\n"
	FixtureStaleReadings  = "ERR DATA-STALE\n"
	FixtureNoDriver       = "ERR DRIVER-NOT-CONNECTED\n"
)
