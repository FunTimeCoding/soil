package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Alert() {
	alerts, _ := common.Alertmanager().MustAlerts(advanced_option.New(), nil)
	f := constant.ExtendedColorFormat.Copy()

	for _, a := range alerts {
		console.Format("Alert: %+v\n", a.Format(f))
	}
}
