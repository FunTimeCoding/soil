package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) formatTeamName(s *option.Format) string {
	var result string

	if a.Team != nil {
		return a.TeamKeyFallback()
	}

	result = NoTeam

	if s.UseColor {
		result = console.Red("%s", result)
	}

	return result
}
