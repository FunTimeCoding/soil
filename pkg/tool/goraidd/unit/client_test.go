package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/generative/constant"
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
	v := service_tester.New(t, cache, output)
	v.SyncLogCache()
	c := mountedClient(t, v, constant.ModelContextTestToken)
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
	v := service_tester.New(t, t.TempDir(), t.TempDir())
	c := mountedClient(t, v, "india-wrong")
	_, e := c.GetReports(t.Context())
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "reports status: 401", e.Error())
	assert.NotNil(t, c.Reachable(t.Context()))
}
