package issue

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Issue) FormatScore(f *option.Format) string {
	s := i.Score()

	if s == 0 {
		return NoScore
	}

	result := fmt.Sprintf("%.1f", s)

	if f.UseColor {
		if i.scoreColor != nil {
			return i.scoreColor(result)
		}

		return console.Yellow("%s", result)
	}

	return result
}
