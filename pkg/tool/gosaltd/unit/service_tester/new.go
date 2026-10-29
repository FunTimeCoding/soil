package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/service"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *service.Service {
	t.Helper()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "gosaltd.db")
	o.Master = func() *salt.Client {
		return salt.New("localhost", 1, "guard", "guard")
	}
	result := service.New(o, logger.New(t.Context()), memory.New())
	t.Cleanup(result.Close)

	return result
}
