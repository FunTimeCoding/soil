package service_tester

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/argocd/mock_client"
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	nextcloud "github.com/funtimecoding/soil/pkg/nextcloud/face"
	usageMock "github.com/funtimecoding/soil/pkg/nextcloud/usage/mock_client"
	metric "github.com/funtimecoding/soil/pkg/prometheus/face"
	prometheusMock "github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/option"
	"path/filepath"
	"testing"
)

func NewOption(
	t *testing.T,
	b *board.Board,
	p *prometheusMock.Client,
	a *mock_client.Client,
	u *usageMock.Client,
) *option.Dashboard {
	t.Helper()
	result := option.New()
	result.Board = b
	result.LitePath = filepath.Join(t.TempDir(), "godashboardd.db")
	result.Issuer = "https://gate.example.org"
	result.ClientIdentifier = "tester"
	result.ClientSecret = "tester-secret"
	result.EncryptionSecret = "tester-encryption-secret"
	result.PublicLocator = "https://dashboard.example.org"
	result.Prometheus = func() metric.Client { return p }

	if a != nil {
		result.Argocd = func() kubernetes.ArgocdClient { return a }
	}

	if u != nil {
		result.Usage = func() nextcloud.UsageClient { return u }
	}

	return result
}
