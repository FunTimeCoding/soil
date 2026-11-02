package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/option"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"testing"
)

func WithMemory(t *testing.T) (*service.Service, *mock_client.Client) {
	t.Helper()
	m := mock_client.New()
	o := option.New()
	o.Inventory = inventory.New()
	o.Memory = func() face.Client { return m }
	result := service.New(o, logger.New(t.Context()), memory.New())
	t.Cleanup(result.Close)
	result.UseIndex(IndexDirectory())

	return result, m
}
