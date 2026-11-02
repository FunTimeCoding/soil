package service_tester

import (
	"github.com/funtimecoding/soil/pkg/directory/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
) *service.Service {
	t.Helper()
	g := peer.New(
		constant.DirectoryPeer,
		"directory.test",
		c,
		logger.New(t.Context()),
		r,
	)
	t.Cleanup(g.Close)

	return service.New(c, g)
}
