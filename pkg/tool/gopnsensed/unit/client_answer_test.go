package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/mock_source"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/missing_host"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/rejecting_host"
	"testing"
)

func TestTheClientReadsTheDaemon(t *testing.T) {
	c := client_tester.New(t, mock_source.New())
	_, e := c.Leases(context.Background(), "")
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, c.Reachable(context.Background()))
}

func TestTheClientAnswersARejectedHostAsValidation(t *testing.T) {
	c := client_tester.New(t, rejecting_host.New())
	_, e := c.AddHost(context.Background(), client.HostRequest{}, true)
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "host rejected: ip: invalid address", e.Error())
}

func TestTheClientAnswersAnUnknownHostAsNotFound(t *testing.T) {
	c := client_tester.New(t, missing_host.New())
	assert.True(
		t,
		not_found.Is(c.DeleteHost(context.Background(), "alfa", true)),
	)
}

func TestTheClientAcceptsANoContentWrite(t *testing.T) {
	c := client_tester.New(t, mock_source.New())
	assert.FatalOnError(
		t,
		c.SetHost(context.Background(), "alfa", client.HostRequest{}, false),
	)
	assert.FatalOnError(t, c.ReconfigureDnsmasq(context.Background()))
}
