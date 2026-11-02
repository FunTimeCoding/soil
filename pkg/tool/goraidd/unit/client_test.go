package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"path/filepath"
	"testing"
)

func TestTheClientListsLogsAndReports(t *testing.T) {
	cache := t.TempDir()
	service_tester.WriteLogCache(
		cache,
		"20261024-201500.zevtc",
		"20261024-203000.zevtc",
	)
	output := t.TempDir()
	system.WriteFile(filepath.Join(output, "hotel.html"), []byte("h"), 0o600)
	o := service_tester.NewOption(t, cache, output, service_tester.NewParser(t))
	service_tester.NewService(t, o, memory.New()).SyncLogCache()
	c := mountedClient(t, o, generativeConstant.ModelContextTestToken)
	logs, e := c.GetLogs(t.Context(), &client.GetLogsParams{})
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, logs.Total)
	reports, f := c.GetReports(t.Context())
	assert.FatalOnError(t, f)
	assert.Count(t, 1, reports)
	assert.String(t, "hotel.html", reports[0].FileName)
	assert.FatalOnError(t, c.Reachable(t.Context()))
}

func TestARefusedTokenIsAnError(t *testing.T) {
	o := service_tester.NewOption(
		t,
		t.TempDir(),
		t.TempDir(),
		service_tester.NewParser(t),
	)
	c := mountedClient(t, o, "india-wrong")
	_, e := c.GetReports(t.Context())
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "reports status: 401", e.Error())
	assert.NotNil(t, c.Reachable(t.Context()))
}
