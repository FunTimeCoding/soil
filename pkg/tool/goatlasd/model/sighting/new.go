package sighting

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/attribution"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"time"
)

func New(
	source string,
	hardwareAddress string,
	address string,
	hostname string,
	reserved bool,
	p *place.Place,
	seenAt time.Time,
) *Sighting {
	return &Sighting{
		Source:          source,
		HardwareAddress: hardwareAddress,
		Address:         address,
		Hostname:        hostname,
		Reserved:        reserved,
		Attribution:     *attribution.New(p, seenAt),
	}
}
