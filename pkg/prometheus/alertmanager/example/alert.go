package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Alert() {
	c := common.Alertmanager()
	defer c.Close()
	p := prometheus.NewEnvironment()
	defer p.Close()
	alerts, _ := c.MustAlerts(
		advanced_option.New(),
		alert_filter.New(),
		p.MustRules(),
	)
	f := constant.ExtendedColorFormat.Copy()

	for _, a := range alerts {
		console.Format("Alert: %+v\n", a.Format(f))
	}
}
