package mock_collector

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
)

func (c *Collector) Collect(
	_ context.Context,
	_ *gazetteer.Gazetteer,
) ([]*placement.Placement, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	return c.placements, nil
}
