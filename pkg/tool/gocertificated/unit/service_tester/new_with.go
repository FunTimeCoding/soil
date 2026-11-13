package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"strconv"
	"testing"
	"time"
)

func NewWith(
	t *testing.T,
	clock func() time.Time,
) (*service.Service, *mock_client.Client) {
	t.Helper()
	d := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(d) })
	f := NewForge(t)
	g := peer.New(
		constant.ForgePeer,
		"",
		f,
		logger.New(t.Context()),
		memory.New(),
	)
	t.Cleanup(g.Close)

	return service.New(
		store.New(d),
		g,
		publish.New(
			f,
			strconv.Itoa(constant.FixtureProject),
			constant.PublishBranch,
			constant.FixtureAuthorityDirectory,
			constant.FixtureClusterAuthority,
			constant.FixtureSecretPath,
		),
		clock,
	), f
}
