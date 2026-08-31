package chart

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/chart/series"
	"maragu.dev/gomponents"
	"time"
)

func (c *Chart) projectionLine(s *series.Series) []gomponents.Node {
	if c.now.IsZero() || len(s.Value) == 0 || !c.now.After(c.start) {
		return nil
	}

	current := s.Value[len(s.Value)-1]
	elapsed := c.now.Sub(c.start).Seconds()
	rate := current / elapsed
	target := c.end
	projected := current + rate*c.end.Sub(c.now).Seconds()

	if projected > c.maximum && rate > 0 {
		remaining := (c.maximum - current) / rate
		target = c.now.Add(time.Duration(remaining * float64(time.Second)))
		projected = c.maximum
	}

	return []gomponents.Node{
		svgLine(
			c.horizontal(c.now),
			c.vertical(current),
			c.horizontal(target),
			c.vertical(projected),
			fmt.Sprintf("chart-projection %s", s.Class),
		),
		svgText(
			c.horizontal(target)+4,
			c.vertical(projected)+3,
			fmt.Sprintf("chart-label chart-projection-label %s", s.Class),
			fmt.Sprintf("%.0f%%", projected),
		),
	}
}
