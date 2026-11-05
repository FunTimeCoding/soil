package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	soilErrors "github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	gomonitord "github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/unit/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
	"time"
)

func TestTheClientClaimsReleasesAndStreams(t *testing.T) {
	s := base.New(t)
	c := gomonitord.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	x, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	events := make(chan []client.Claim, 8)
	ended := make(chan error, 1)
	b, e := c.Listen(x)
	assert.FatalOnError(t, e)
	defer soilErrors.LogClose(b)
	go func() {
		ended <- gomonitord.ReadClaims(
			b,
			func(v []client.Claim) { events <- v },
		)
	}()
	assert.Count(t, 0, nextEvent(t, events))
	put, f := c.PutClaim(x, "jira-ABC-1", "alfa@host.example")
	assert.FatalOnError(t, f)
	assert.String(t, "alfa@host.example", put.Owner)
	claimed := nextEvent(t, events)
	assert.Count(t, 1, claimed)
	assert.String(t, "alfa@host.example", claimed[0].Owner)
	assert.FatalOnError(t, c.DeleteClaim(x, "jira-ABC-1", "alfa@host.example"))
	assert.Count(t, 0, nextEvent(t, events))
	cancel()
	assert.True(t, errors.Is(streamEnd(t, ended), context.Canceled))
}

func TestListenReportsAnUnreachableDaemon(t *testing.T) {
	c := gomonitord.New(
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
	c := gomonitord.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		"bogus",
	)
	defer c.Close()
	_, e := c.Listen(context.Background())
	assert.StringContains(t, "stream status: 401", e.Error())
}
