package unit

import (
	"context"
	standard "errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/unit/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
	"time"
)

func TestTheClientClaimsReleasesAndStreams(t *testing.T) {
	s := base.New(t)
	c := client.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	x, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	events := make(chan []generated.Claim, 8)
	ended := make(chan error, 1)
	b, e := c.Listen(x)
	assert.FatalOnError(t, e)
	defer errors.LogClose(b)
	go func() {
		ended <- client.ReadClaims(b, func(v []generated.Claim) { events <- v })
	}()
	assert.Count(t, 0, nextEvent(t, events))
	assert.FatalOnError(t, c.ClaimItem("jira-ABC-1", "alfa@host.example"))
	claimed := nextEvent(t, events)
	assert.Count(t, 1, claimed)
	assert.String(t, "alfa@host.example", claimed[0].Owner)
	assert.FatalOnError(t, c.Release("jira-ABC-1", "alfa@host.example"))
	assert.Count(t, 0, nextEvent(t, events))
	cancel()
	assert.True(t, standard.Is(streamEnd(t, ended), context.Canceled))
}

func TestListenReportsAnUnreachableDaemon(t *testing.T) {
	c := client.New(
		locator.New(webConstant.Localhost).Port(1).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	x, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)
	_, e := c.Listen(x)
	assert.True(t, connection.Classify(e) != nil)
}

func TestListenRefusesAnUnknownToken(t *testing.T) {
	s := base.New(t)
	c := client.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		"bogus",
	)
	defer c.Close()
	_, e := c.Listen(context.Background())
	assert.StringContains(t, "stream status: 401", e.Error())
}
