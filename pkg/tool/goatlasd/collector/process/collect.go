package process

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"time"
)

func (c *Collector) Collect(
	q context.Context,
	s *gazetteer.Gazetteer,
) ([]*placement.Placement, error) {
	processes, e := c.process.ListProcessesWithResponse(q)

	if e != nil {
		return nil, e
	}

	if processes.JSON200 == nil {
		return nil, unexpected.Format(
			constant.ProcessStatusFormat,
			processes.HTTPResponse.StatusCode,
		)
	}

	where, okay := s.Resolve(c.place)

	if !okay {
		return nil, unexpected.Format(constant.PlaceUnknownFormat, c.place)
	}

	var result []*placement.Placement
	now := time.Now()

	for _, p := range *processes.JSON200 {
		if !p.Running {
			continue
		}

		result = append(
			result,
			placement.New(
				constant.SourceProcess,
				constant.KindService,
				"",
				p.Name,
				where,
				now,
			),
		)
	}

	return result, nil
}
