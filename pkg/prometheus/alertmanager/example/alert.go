package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Alert() {
	alerts, _ := common.Alertmanager().MustAlerts(advanced_option.New(), nil)
	f := option.ExtendedColor.Copy()

	for _, a := range alerts {
		fmt.Printf("Alert: %+v\n", a.Format(f))
	}
}
