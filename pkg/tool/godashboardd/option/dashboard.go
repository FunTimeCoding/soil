package option

import (
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	nextcloud "github.com/funtimecoding/soil/pkg/nextcloud/face"
	metric "github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
)

type Dashboard struct {
	Address          string
	Board            *board.Board
	PostgresLocator  string
	LitePath         string
	Issuer           string
	ClientIdentifier string
	ClientSecret     string
	EncryptionSecret string
	PublicLocator    string
	ServiceTokens    []string
	Prometheus       func() metric.Client
	Argocd           func() kubernetes.ArgocdClient
	Usage            func() nextcloud.UsageClient
}
