package common

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/alert_enricher"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/field_changer"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/name_filter"
	prometheusConstant "github.com/funtimecoding/soil/pkg/prometheus/constant"
)

func Alertmanager() *alertmanager.Client {
	return alertmanager.NewEnvironment().Set(
		alert_enricher.New().Add(
			prometheusConstant.KubernetesCronJobFailed,
			prometheusConstant.EntityJob,
			prometheusConstant.CategoryFail,
		).Add(
			prometheusConstant.Watchdog,
			prometheusConstant.EntityService,
			prometheusConstant.CategoryOkay,
		),
		field_changer.New(),
		name_filter.New(true),
		label_filter.New(true),
	)
}
