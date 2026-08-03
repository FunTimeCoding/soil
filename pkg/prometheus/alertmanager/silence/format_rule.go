package silence

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
)

func (s *Silence) formatRule(f *option.Format) string {
	var result string

	if s.Rule != constant.UnknownRule {
		result = s.Rule
	} else {
		result = s.Match
	}

	if f.UseColor {
		result = consoleConstant.Cyan("%s", result)
	}

	return result
}
