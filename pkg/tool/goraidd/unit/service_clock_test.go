package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	gw2Constant "github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCleanupCutoffRunsOnServiceClock(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithClock(
			func() time.Time { return time.Now().AddDate(0, 0, 30) },
		),
	)
	recent := filepath.Join(o.Elite, "recent.json")
	assert.FatalOnError(t, os.WriteFile(recent, nil, 0o600))
	day := time.Now().AddDate(0, 0, -1)
	assert.FatalOnError(t, os.Chtimes(recent, day, day))
	s := o.Service()
	assert.Count(t, 0, s.Cleanup(t.Context()))
	_, e := os.Stat(recent)
	assert.True(t, os.IsNotExist(e))
}

func TestEnrichCutoffRunsOnServiceClock(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithClock(
			func() time.Time { return time.Now().AddDate(0, 0, 30) },
		),
	)
	name := join.Empty(
		time.Now().AddDate(0, 0, -1).Format(constant.FileDateLayout),
		"_clocked",
		gw2Constant.DetailedWvWKillSuffix,
	)
	assert.FatalOnError(
		t,
		os.WriteFile(filepath.Join(o.Elite, name), []byte("{"), 0o600),
	)
	s := o.Service()
	assert.Count(t, 0, s.EnrichExisting(t.Context()))
}
