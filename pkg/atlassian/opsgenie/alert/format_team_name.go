package alert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) formatTeamName(s *option.Format) string {
	var result string

	if a.Team != nil {
		return a.TeamKeyFallback()
	}

	result = constant.OpsgenieNoTeam

	if s.UseColor {
		result = consoleConstant.Red("%s", result)
	}

	return result
}
