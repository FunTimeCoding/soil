package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/hold"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/unit/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
	"time"
)

func TestAWindowReceivesClaimsThroughTheHold(t *testing.T) {
	s := base.New(t)
	g := nextClaimMessage(
		t,
		s.Port,
		generativeConstant.ModelContextTestToken,
		memory.New(),
	)
	assert.Nil(t, g.Error)
}

func TestAClosingWindowReleasesANoticeNobodyReads(t *testing.T) {
	s := base.New(t)
	m := client.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		"bogus",
	)
	defer m.Close()
	updates := make(chan claim.Message, 1)
	updates <- claim.Message{}
	c := claim.New(m, updates)
	h := hold.New(
		"gomonitord",
		c,
		logger.New(t.Context()),
		memory.New(),
	).WithReport(c.Crossed, c.Refused)
	h.Start()
	closed := make(chan struct{})
	go func() {
		h.Close()
		close(closed)
	}()

	select {
	case <-closed:
	case <-time.After(2 * time.Second):
		t.Fatal("close waited on a notice nobody reads")
	}
}

func TestAWindowShowsARefusedClaimStream(t *testing.T) {
	s := base.New(t)
	r := memory.New()
	g := nextClaimMessage(t, s.Port, "bogus", r)
	assert.StringContains(t, "stream status: 401", g.Error.Error())
	assert.Count(t, 0, r.Events())
}
