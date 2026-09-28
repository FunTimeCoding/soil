package mock_collector

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"

func (c *Collector) Add(p *placement.Placement) {
	c.placements = append(c.placements, p)
}
