package collector_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"time"
)

func NewSighting(
	address string,
	hostname string,
	p *place.Place,
) *sighting.Sighting {
	return sighting.New(
		constant.SourceLease,
		constant.FixtureHardwareAddress,
		address,
		hostname,
		false,
		p,
		time.Now(),
	)
}
