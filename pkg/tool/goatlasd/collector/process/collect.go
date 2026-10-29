package process

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
	"time"
)

func (c *Collector) Collect(
	q context.Context,
	s *gazetteer.Gazetteer,
) ([]*placement.Placement, error) {
	var processes []client.Process
	e := c.guard.Attempt(
		q,
		func(x context.Context) error {
			var f error
			processes, f = c.process.Processes(x)

			return f
		},
	)

	if e != nil {
		return nil, e
	}

	where, okay := s.Resolve(c.place)

	if !okay {
		return nil, unexpected.Format(constant.PlaceUnknownFormat, c.place)
	}

	var result []*placement.Placement
	now := time.Now()

	for _, p := range processes {
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
