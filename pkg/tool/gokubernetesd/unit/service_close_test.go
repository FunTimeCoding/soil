package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/option"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"path/filepath"
	"testing"
)

func TestTheServiceClosesTheDatabaseItOpened(t *testing.T) {
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "gokubernetesd.db")
	o.Clusters = func() map[string]*cluster.Cluster { return nil }
	s := service.New(o, logger.New(t.Context()), memory.New())
	_, e := s.MuteRules()
	assert.FatalOnError(t, e)
	s.Close()
	_, f := s.MuteRules()
	assert.FatalNotNil(t, f)
	assert.StringContains(t, "database is closed", f.Error())
}
