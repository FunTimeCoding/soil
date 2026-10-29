package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/unit/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"testing"
)

func TestTheClientClassifiesRefusals(t *testing.T) {
	c := mattermost_client_tester.New(t, statusUpstream).Client
	_, e := c.FindPost("forbidden")
	assert.True(t, forbidden.Is(e))
	_, f := c.FindPost("missing")
	assert.True(t, not_found.Is(f))
	_, g := c.FindPost("proxied")
	assert.True(t, unavailable.Is(g))
}

func TestAnUnauthorizedTokenStaysAFault(t *testing.T) {
	c := mattermost_client_tester.New(t, statusUpstream).Client
	_, e := c.FindPost("unauthorized")
	assert.True(t, classify.Reportable(e))
	_, okay := e.(*detail_error.Detail)
	assert.True(t, okay)
}

func TestAnUnknownStatusKeepsItsDetail(t *testing.T) {
	c := mattermost_client_tester.New(t, statusUpstream).Client
	_, e := c.FindPost("teapot")
	d, okay := e.(*detail_error.Detail)
	assert.True(t, okay)
	assert.StringContains(t, "418", d.Status)
}

func TestATransportFailureKeepsItsCause(t *testing.T) {
	r := mattermost_client_tester.New(t, statusUpstream)
	r.Stop()
	_, e := r.Client.FindPost("missing")
	assert.NotNil(t, connection.Classify(e))
}
