package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/directory/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/option"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	"testing"
)

func Shared(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
) *service.Service {
	t.Helper()
	o := option.New()
	o.Host = "directory.test"
	o.Issuer = "https://gate.example.org"
	o.ClientIdentifier = "tester"
	o.ClientSecret = "tester-secret"
	o.EncryptionSecret = "tester-encryption-secret"
	o.PublicLocator = "https://directory.example.org"
	o.Source = func() face.Client { return c }
	result := service.New(o, logger.New(context.Background()), r)
	t.Cleanup(result.Close)

	return result
}
