package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/habitica/face"
	"github.com/funtimecoding/soil/pkg/habitica/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/option"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/service"
	"testing"
)

func Shared(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
) *service.Service {
	t.Helper()
	o := option.New()
	o.Host = "habitica.test"
	o.Source = func() face.Client { return c }
	result := service.New(o, logger.New(context.Background()), r)
	t.Cleanup(result.Close)

	return result
}
