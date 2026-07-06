package alert

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (a *Alert) buildLink(host string) string {
	var query string

	if i := a.Detail(prometheus.InstanceLabel); i != "" {
		query = fmt.Sprintf(
			"{%s=\"%s\",%s=\"%s\"}",
			constant.AlertnameLabel,
			a.Name,
			prometheus.InstanceLabel,
			i,
		)
	} else {
		query = fmt.Sprintf(
			"{%s=\"%s\"}",
			constant.AlertnameLabel,
			a.Name,
		)
	}

	return locator.New(host).Trail().Fragment(
		constant.Alerts,
	).FragmentSet("filter", query).String()
}
