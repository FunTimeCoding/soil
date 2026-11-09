package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/rejecting_host"
	"testing"
)

func TestRejectedHostIsAnsweredAsBadRequest(t *testing.T) {
	m := memory.New()
	c := client_tester.NewWith(t, rejecting_host.New(), m)
	_, e := c.AddHost(context.Background(), client.HostRequest{}, true)
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "host rejected: ip: invalid address", e.Error())
	assert.Count(t, 0, m.Events())
}
