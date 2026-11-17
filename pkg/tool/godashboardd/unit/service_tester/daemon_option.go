package service_tester

import (
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	nextcloud "github.com/funtimecoding/soil/pkg/nextcloud/face"
	metric "github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/option"
)

func (o *Tester) daemonOption() *option.Dashboard {
	result := option.New()
	result.Board = o.board
	result.LitePath = o.litePath
	result.Prometheus = func() metric.PrometheusClient { return o.Prometheus }

	if o.Argocd != nil {
		result.Argocd = func() kubernetes.ArgocdClient { return o.Argocd }
	}

	if o.Usage != nil {
		result.Usage = func() nextcloud.UsageClient { return o.Usage }
	}

	return result
}
