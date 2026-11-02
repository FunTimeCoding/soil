package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/option"
	"path/filepath"
	"testing"
)

func NewOption(
	t *testing.T,
	clusters map[string]*cluster.Cluster,
) *option.Server {
	t.Helper()
	result := option.New()
	result.LitePath = filepath.Join(t.TempDir(), "gokubernetesd.db")
	result.Clusters = func() map[string]*cluster.Cluster { return clusters }

	return result
}
