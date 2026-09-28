package lease

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
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
	leases, e := c.opnsense.ListLeasesWithResponse(
		q,
		&client.ListLeasesParams{},
	)

	if e != nil {
		return nil, e
	}

	if leases.JSON200 == nil {
		return nil, unexpected.Format(
			constant.LeaseStatusFormat,
			leases.HTTPResponse.StatusCode,
		)
	}

	var result []*sighting.Sighting
	now := time.Now()

	for _, l := range *leases.JSON200 {
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
