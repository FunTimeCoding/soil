package main

import (
	grafana "github.com/funtimecoding/go-library/pkg/grafana/example"
	alertmanager "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/example"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/example/notify"
	prometheus "github.com/funtimecoding/go-library/pkg/prometheus/example"
	loki "github.com/funtimecoding/go-library/pkg/prometheus/loki/example"
)

func main() {
	loki.QueryRange()
	//prometheus.Target()
	//prometheus.Series()
	if false {
		alertmanager.Alert()
		alertmanager.Create()
		alertmanager.DeleteSilence()
		notify.Notify()
		alertmanager.SetSilence()
		alertmanager.Silence()
		alertmanager.Status()
		grafana.Read()
		loki.Label()
		loki.Official()
		loki.Query()
		loki.Series()
		loki.Statistic()
		loki.Write()
		prometheus.Label()
		prometheus.LabelName()
		prometheus.Meta()
		prometheus.Metric()
		prometheus.Query()
		prometheus.Rule()
		prometheus.Status()
	}
}
