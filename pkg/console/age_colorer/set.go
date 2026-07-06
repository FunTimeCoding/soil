package age_colorer

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/math/in_range"
)

func (c *Colorer) Set(i face.AgeColorable) {
	value := i.Age().Hours() / c.largest

	for _, m := range c.mapping {
		if in_range.LeftOpen(value, m.Range) {
			if f, okay := c.assignments[m.Value]; okay {
				i.SetAgeColor(f)
			}

			break
		}
	}

	if first := c.mapping[0]; value <= first.Range.L {
		i.SetAgeColor(console.Green)
	}

	if last := c.mapping[len(c.mapping)-1]; value >= last.Range.R {
		i.SetAgeColor(console.Red)
	}

	if i.AgeColor() == nil {
		panic("unable to determine color")
	}
}
