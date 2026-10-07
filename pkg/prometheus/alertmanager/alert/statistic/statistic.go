package statistic

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/statistic/count"

type Statistic struct {
	Total    int
	Relevant int
	Severity count.Severity
	State    count.State
	Group    count.Group
}
