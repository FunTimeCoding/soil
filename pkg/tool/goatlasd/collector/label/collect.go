package label

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"time"
)

func (c *Collector) Collect(
	q context.Context,
	_ *gazetteer.Gazetteer,
) ([]*placement.Placement, error) {
	now := time.Now()
	result, e := c.devicePlacements(q, now)

	if e != nil {
		return nil, e
	}

	machines, f := c.virtualPlacements(q, now)

	if f != nil {
		return nil, f
	}

	return append(result, machines...), nil
}
