package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/query_option"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/unit/web_service_tester"
	"net/http"
	"testing"
	"time"
)

func TestDaemonClosesItsDatabaseOnce(t *testing.T) {
	d := daemon.New(base.NewOption(t), logger.New(t.Context()), memory.New())
	d.Close()
	d.Close()
}

func TestServiceRecordsItsOwnUse(t *testing.T) {
	o := base.NewOption(t)
	m := lite.New(logger.New(t.Context()), o.LitePath)
	defer relational.CloseMapper(m)
	migrate.AutoMigrate(m)
	v := service.New(store.New(m), time.Now, memory.New())
	e := record.NewBaseline(
		"GetSummary",
		constant.SurfaceWebService,
		"tester",
		constant.OutcomeSuccess,
	)
	e.Detail = map[string]string{"operation": "GetSummary"}
	v.Record(e)
	events, f := v.Query(query_option.New())
	assert.FatalOnError(t, f)
	assert.Count(t, 1, events)
	assert.String(t, "GetSummary", events[0].Tool)
	assert.NotNil(t, events[0].Detail)
}

func TestFailingStoreIsAnsweredAndReported(t *testing.T) {
	o := web_service_tester.New(t)
	o.Daemon.Close()
	r, e := o.Client.GetSummaryWithResponse(
		context.Background(),
		&client.GetSummaryParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, r.StatusCode())
	assert.NotNil(t, r.JSON500)
	assert.Count(t, 2, o.Reporter.Events())
}
