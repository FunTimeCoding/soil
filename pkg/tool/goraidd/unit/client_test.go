package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"path/filepath"
	"testing"
)

func TestClientListsLogsAndReports(t *testing.T) {
	s := base.New(t)
	service_tester.WriteLogCache(
		s.LogCache,
		"20261024-201500.zevtc",
		"20261024-203000.zevtc",
	)
	system.WriteFile(filepath.Join(s.Output, "hotel.html"), []byte("h"), 0o600)
	s.Service().SyncLogCache(t.Context())
	c := s.Client(generativeConstant.ModelContextTestToken)
	logs, e := c.GetLogs(t.Context(), &client.GetLogsParams{})
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, logs.Total)
	reports, f := c.GetReports(t.Context())
	assert.FatalOnError(t, f)
	assert.Count(t, 1, reports)
	assert.String(t, "hotel.html", reports[0].FileName)
	assert.FatalOnError(t, c.Reachable(t.Context()))
}

func TestRefusedTokenIsError(t *testing.T) {
	c := base.New(t).Client("india-wrong")
	_, e := c.GetReports(t.Context())
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "reports status: 401", e.Error())
	assert.NotNil(t, c.Reachable(t.Context()))
}
