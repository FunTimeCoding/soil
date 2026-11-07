package option

import (
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	nextcloud "github.com/funtimecoding/soil/pkg/nextcloud/face"
	metric "github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Dashboard struct {
	option.Web
	Board           *board.Board
	PostgresLocator string
	LitePath        string
	Prometheus      func() metric.PrometheusClient
	Argocd          func() kubernetes.ArgocdClient
	Usage           func() nextcloud.UsageClient
}
