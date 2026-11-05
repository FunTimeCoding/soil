package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/opnsense/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/service"
	"testing"
)

func New(
	t *testing.T,
	c face.OpnsenseClient,
) *service.Service {
	t.Helper()
	g := peer.New(
		constant.OpnsensePeer,
		c.Address(),
		c,
		logger.New(t.Context()),
		memory.New(),
	)
	t.Cleanup(g.Close)

	return service.New(c, g)
}
