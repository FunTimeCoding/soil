package lease

import "time"

type Lease struct {
	Address          string
	HardwareAddress  string
	Hostname         string
	ClientIdentifier string
	Expire           time.Time
	Interface        string
	Reserved         bool
}
