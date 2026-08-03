package issue

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Issue) FormatScore(f *option.Format) string {
	s := i.Score()

	if s == 0 {
		return constant.JiraNoScore
	}

	result := fmt.Sprintf("%.1f", s)

	if f.UseColor {
		if i.scoreColor != nil {
			return i.scoreColor(result)
		}

		return consoleConstant.Yellow("%s", result)
	}

	return result
}
