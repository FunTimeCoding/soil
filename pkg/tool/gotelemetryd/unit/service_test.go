package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/query_option"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/unit/web_service_tester"
	"net/http"
	"testing"
)

func TestTheServiceClosesItsDatabaseOnce(t *testing.T) {
	v := service.New(base.NewOption(t), logger.New(t.Context()), memory.New())
	v.Close()
	v.Close()
	_, e := v.Query(query_option.New())
	assert.Error(t, e)
}

func TestTheServiceRecordsItsOwnUse(t *testing.T) {
	b := base.New(t)
	e := record.NewBaseline(
		"GetSummary",
		constant.SurfaceWebService,
		"tester",
		constant.OutcomeSuccess,
	)
	e.Detail = map[string]string{"operation": "GetSummary"}
	b.Service.Record(e)
	events, f := b.Service.Query(query_option.New())
	assert.FatalOnError(t, f)
	assert.Count(t, 1, events)
	assert.String(t, "GetSummary", events[0].Tool)
	assert.NotNil(t, events[0].Detail)
}

func TestAFailingStoreIsAnsweredAndReported(t *testing.T) {
	o := web_service_tester.New(t)
	o.Service.Close()
	r, e := o.Client.GetSummaryWithResponse(
		context.Background(),
		&client.GetSummaryParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, r.StatusCode())
	assert.NotNil(t, r.JSON500)
	assert.Count(t, 1, o.Reporter.Events())
}
