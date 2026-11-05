package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/iterm/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/service"
	"testing"
	"time"
)

func NewPaced(
	t *testing.T,
	c *mock_client.Client,
	keyInterval time.Duration,
) *service.Service {
	t.Helper()
	g := peer.New(
		constant.ItermPeer,
		c.Address(),
		c,
		logger.New(t.Context()),
		memory.New(),
	)
	t.Cleanup(g.Close)

	return service.New(c, g, keyInterval)
}
