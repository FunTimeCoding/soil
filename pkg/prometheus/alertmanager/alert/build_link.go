package alert

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (a *Alert) buildLink(link *locator.Locator) string {
	var query string

	if i := a.Detail(constant.InstanceLabel); i != "" {
		query = fmt.Sprintf(
			"{%s=\"%s\",%s=\"%s\"}",
			constant.AlertnameLabel,
			a.Name,
			constant.InstanceLabel,
			i,
		)
	} else {
		query = fmt.Sprintf("{%s=\"%s\"}", constant.AlertnameLabel, a.Name)
	}

	return link.Copy().Trail().Fragment(constant.Alerts).FragmentSet(
		"filter",
		query,
	).String()
}
