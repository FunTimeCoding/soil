package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/service_tester"
	"sync/atomic"
	"testing"
	"time"
)

func TestExpiryHorizonRunsOnServiceClock(t *testing.T) {
	var now atomic.Int64
	now.Store(time.Now().UnixNano())
	s := base.New(
		t,
		service_tester.WithClock(
			func() time.Time { return time.Unix(0, now.Load()) },
		),
	)
	s.Submit(constant.CreateAuthorityPath, base.RootValues())
	d := relational.Open(logger.New(t.Context()), "", s.LitePath)
	t.Cleanup(func() { relational.CloseMapper(d) })
	root := record.Record{}
	assert.FatalOnError(t, d.First(&root).Error)
	now.Store(root.End.AddDate(0, 0, -10).UnixNano())
	assert.StringContains(
		t,
		"1 expiring within a month",
		s.Page(constant.DashboardPath),
	)
}
