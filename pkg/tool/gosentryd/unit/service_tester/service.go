package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/face"
	"github.com/funtimecoding/soil/pkg/errors/sentry/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/service"
	"testing"
)

func Service(
	t *testing.T,
	c *mock_client.Client,
) *service.Service {
	t.Helper()
	o := option.New()
	o.Organization = "test"
	o.Host = "sentry.test"
	o.Source = func() face.Client { return c }
	result := service.New(o, logger.New(context.Background()), memory.New())
	t.Cleanup(result.Close)

	return result
}
