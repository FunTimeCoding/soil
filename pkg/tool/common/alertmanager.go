package common

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/alert_enricher"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/field_changer"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/name_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
)

func Alertmanager() *alertmanager.Client {
	return alertmanager.NewEnvironment().Set(
		alert_enricher.New().Add(
			constant.KubernetesCronJobFailed,
			constant.EntityJob,
			constant.CategoryFail,
		).Add(
			constant.Watchdog,
			constant.EntityService,
			constant.CategoryOkay,
		),
		field_changer.New(),
		name_filter.New(true),
		label_filter.New(true),
	)
}
