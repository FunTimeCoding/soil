package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/habitica/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/service"
	"testing"
)

func Shared(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
) *service.Service {
	t.Helper()
	g := peer.New(
		constant.HabiticaPeer,
		"habitica.test",
		c,
		logger.New(context.Background()),
		r,
	)
	t.Cleanup(g.Close)

	return service.New(c, g)
}
