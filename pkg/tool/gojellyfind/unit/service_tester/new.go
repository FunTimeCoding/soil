package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/jellyfin/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/constant"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/service"
	"testing"
)

func New(
	t *testing.T,
	c face.JellyfinClient,
) *service.Service {
	t.Helper()
	g := peer.New(
		constant.JellyfinPeer,
		c.Address(),
		c,
		logger.New(t.Context()),
		memory.New(),
	)
	t.Cleanup(g.Close)

	return service.New(c, g)
}
