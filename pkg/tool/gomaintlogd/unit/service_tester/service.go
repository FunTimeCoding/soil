package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/service"
	"path/filepath"
	"testing"
)

func Service(t *testing.T) *service.Service {
	t.Helper()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "maintenance.db")
	result := service.New(o, logger.New(t.Context()))
	t.Cleanup(result.Close)

	return result
}
