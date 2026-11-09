package option

import "time"

func New() *Telemetry {
	return &Telemetry{Clock: time.Now}
}
