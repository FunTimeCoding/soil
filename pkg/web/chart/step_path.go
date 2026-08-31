package chart

import (
	"github.com/funtimecoding/soil/pkg/system/writer"
	"github.com/funtimecoding/soil/pkg/web/chart/series"
	"strings"
)

func (c *Chart) stepPath(s *series.Series) string {
	var b strings.Builder

	for i, at := range s.Times {
		x := c.horizontal(at)
		y := c.vertical(s.Value[i])

		if i == 0 {
			writer.Print(&b, "M %.1f %.1f", x, y)

			continue
		}

		writer.Print(&b, " H %.1f V %.1f", x, y)
	}

	if !c.now.IsZero() && len(s.Times) > 0 {
		writer.Print(&b, " H %.1f", c.horizontal(c.now))
	}

	return b.String()
}
