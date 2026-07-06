package score_colorer

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/math/in_range"
)

func (c *Colorer) Set(i face.ScoreColorable) {
	if c.largest == 0 {
		i.SetScoreColor(console.Green)

		return
	}

	value := i.Score() / c.largest

	for _, m := range c.mapping {
		if in_range.LeftOpen(value, m.Range) {
			if f, okay := c.assignments[m.Value]; okay {
				i.SetScoreColor(f)
			}

			break
		}
	}

	if first := c.mapping[0]; value <= first.Range.L {
		i.SetScoreColor(console.Green)
	}

	if last := c.mapping[len(c.mapping)-1]; value >= last.Range.R {
		i.SetScoreColor(console.Red)
	}

	if i.ScoreColor() == nil {
		panic("unable to determine color")
	}
}
