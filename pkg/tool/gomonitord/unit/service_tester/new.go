package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/option"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/service"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *service.Service {
	t.Helper()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "gomonitord.db")
	result := service.New(o, logger.New(t.Context()))
	t.Cleanup(result.Close)

	return result
}
