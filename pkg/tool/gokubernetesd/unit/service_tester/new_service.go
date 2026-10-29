package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/option"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"path/filepath"
	"testing"
)

func NewService(
	t *testing.T,
	clusters map[string]*cluster.Cluster,
) *service.Service {
	t.Helper()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "gokubernetesd.db")
	o.Clusters = func() map[string]*cluster.Cluster { return clusters }
	result := service.New(o, logger.New(t.Context()), memory.New())
	t.Cleanup(result.Close)

	return result
}
