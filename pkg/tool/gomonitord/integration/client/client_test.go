package client

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/integration/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
	"time"
)

func TestTheClientClaimsReleasesAndStreams(t *testing.T) {
	s := base.New(t)
	c := client.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		constant.ModelContextTestToken,
	)
	x, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	events := make(chan []generated.Claim, 8)
	ended := make(chan error, 1)
	go func() {
		ended <- c.Stream(x, func(v []generated.Claim) { events <- v })
	}()
	assert.Count(t, 0, nextEvent(t, events))
	assert.FatalOnError(t, c.ClaimItem("jira-ABC-1", "alfa@host.example"))
	claimed := nextEvent(t, events)
	assert.Count(t, 1, claimed)
	assert.String(t, "alfa@host.example", claimed[0].Owner)
	assert.FatalOnError(t, c.Release("jira-ABC-1", "alfa@host.example"))
	assert.Count(t, 0, nextEvent(t, events))
	cancel()
	assert.True(t, errors.Is(streamEnd(t, ended), context.Canceled))
}

func TestStreamReportsAnUnreachableDaemon(t *testing.T) {
	c := client.New(
		locator.New(webConstant.Localhost).Port(1).Insecure(),
		constant.ModelContextTestToken,
	)
	x, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)
	assert.True(t, c.Stream(x, func([]generated.Claim) {}) != nil)
}
