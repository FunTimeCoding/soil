package monitor_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim/message"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/hold"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"time"
)

func (o *Tester) ClaimMessage(
	port int,
	token string,
	r *memory.Memory,
) *message.Message {
	o.t.Helper()
	m := client.New(
		locator.New(webConstant.Localhost).Port(port).Insecure(),
		token,
	)
	defer m.Close()
	updates := make(chan message.Message, 1)
	c := claim.New(m, updates)
	h := hold.New(
		"gomonitord",
		c,
		logger.New(context.Background()),
		r,
	).WithReport(c.Crossed, c.Refused)
	h.Start()
	defer h.Close()

	select {
	case g := <-updates:
		return &g
	case <-time.After(2 * time.Second):
		o.t.Fatal("no claim message")

		return nil
	}
}
