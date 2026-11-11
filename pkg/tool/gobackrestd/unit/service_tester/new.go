package service_tester

import (
	"github.com/funtimecoding/soil/pkg/backrest/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/service"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
) *service.Service {
	t.Helper()
	g := peer.New(
		constant.BackrestPeer,
		"pgbackrest",
		c,
		logger.New(t.Context()),
		memory.New(),
	)
	t.Cleanup(g.Close)

	return service.New(c, g)
}
