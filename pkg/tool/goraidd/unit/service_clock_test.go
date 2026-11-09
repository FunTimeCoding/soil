package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
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
	o := service_tester.NewOption(
		t,
		t.TempDir(),
		t.TempDir(),
		service_tester.NewParser(t),
	)
	o.Clock = func() time.Time { return time.Now().AddDate(0, 0, 30) }
	recent := filepath.Join(o.ElitePath, "recent.json")
	assert.FatalOnError(t, os.WriteFile(recent, nil, 0o600))
	day := time.Now().AddDate(0, 0, -1)
	assert.FatalOnError(t, os.Chtimes(recent, day, day))
	s := service_tester.NewService(t, o, memory.New())
	assert.Count(t, 0, s.Cleanup(t.Context()))
	_, e := os.Stat(recent)
	assert.True(t, os.IsNotExist(e))
}

func TestEnrichCutoffRunsOnServiceClock(t *testing.T) {
	o := service_tester.NewOption(
		t,
		t.TempDir(),
		t.TempDir(),
		service_tester.NewParser(t),
	)
	o.Clock = func() time.Time { return time.Now().AddDate(0, 0, 30) }
	name := join.Empty(
		time.Now().AddDate(0, 0, -1).Format(constant.FileDateLayout),
		"_clocked",
		gw2Constant.DetailedWvWKillSuffix,
	)
	assert.FatalOnError(
		t,
		os.WriteFile(filepath.Join(o.ElitePath, name), []byte("{"), 0o600),
	)
	s := service_tester.NewService(t, o, memory.New())
	assert.Count(t, 0, s.EnrichExisting(t.Context()))
}
