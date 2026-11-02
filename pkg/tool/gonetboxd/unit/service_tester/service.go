package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
	"path/filepath"
	"testing"
)

func Service(
	t *testing.T,
	c face.NetBoxClient,
) *service.Service {
	t.Helper()
	d := relational.Open(
		logger.New(t.Context()),
		"",
		filepath.Join(t.TempDir(), "gonetboxd.db"),
	)
	migrate.AutoMigrate(d)
	t.Cleanup(func() { relational.CloseMapper(d) })

	return service.New(store.New(d), c)
}
