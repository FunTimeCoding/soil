package alert

import "github.com/funtimecoding/soil/pkg/atlassian/constant"

func (a *Alert) TeamKeyFallback() string {
	if a.TeamKey == constant.OpsgenieNoKey {
		return a.Team.Name
	}

	return a.TeamKey
}
