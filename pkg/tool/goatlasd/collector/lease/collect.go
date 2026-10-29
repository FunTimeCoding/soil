package lease

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"time"
)

func (c *Collector) Collect(
	q context.Context,
	s *gazetteer.Gazetteer,
) ([]*sighting.Sighting, error) {
	var leases []client.Lease
	e := c.guard.Attempt(
		q,
		func(x context.Context) error {
			var f error
			leases, f = c.opnsense.Leases(x, "")

			return f
		},
	)

	if e != nil {
		return nil, e
	}

	var result []*sighting.Sighting
	now := time.Now()

	for _, l := range leases {
		where, _ := s.ResolveHost(l.Address, l.HardwareAddress)
		result = append(
			result,
			sighting.New(
				constant.SourceLease,
				l.HardwareAddress,
				l.Address,
				l.Hostname,
				l.Reserved,
				where,
				now,
			),
		)
	}

	return result, nil
}
