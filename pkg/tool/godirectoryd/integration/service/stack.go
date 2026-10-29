package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/option"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	"testing"
)

func stack(t *testing.T) *service.Service {
	t.Helper()
	o := option.New()
	o.Host = environment.Required(constant.HostEnvironment)
	o.Issuer = "https://gate.example.org"
	o.ClientIdentifier = "tester"
	o.ClientSecret = "tester-secret"
	o.EncryptionSecret = "tester-encryption-secret"
	o.PublicLocator = "https://directory.example.org"
	result := service.New(o, logger.New(context.Background()), memory.New())
	t.Cleanup(result.Close)

	return result
}
