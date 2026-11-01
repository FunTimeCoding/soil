package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/hold"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
	"time"
)

func nextClaimMessage(
	t *testing.T,
	port int,
	token string,
	r *memory.Memory,
) *claim.Message {
	t.Helper()
	m := client.New(
		locator.New(webConstant.Localhost).Port(port).Insecure(),
		token,
	)
	defer m.Close()
	updates := make(chan claim.Message, 1)
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
		t.Fatal("no claim message")

		return nil
	}
}
