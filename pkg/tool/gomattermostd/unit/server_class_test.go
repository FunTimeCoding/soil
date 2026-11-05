package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"net/http"
	"testing"
	"time"
)

func TestAMissingPostIsAnsweredNotFound(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	c, r := servedDaemon(t, m)
	result, e := c.GetThreadWithResponse(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, result.StatusCode())
	assert.FatalNotNil(t, result.JSON404)
	assert.String(t, "post not found: alfa", result.JSON404.Error)
	assert.Count(t, 0, r.Events())
}

func TestARefusedPostIsAnsweredForbidden(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	m.Failure = forbidden.Format("no access to alfa")
	c, r := servedDaemon(t, m)
	result, e := c.GetPostReactionsWithResponse(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusForbidden, result.StatusCode())
	assert.FatalNotNil(t, result.JSON403)
	assert.String(t, "no access to alfa", result.JSON403.Error)
	assert.Count(t, 0, r.Events())
}

func TestAnUnavailableMattermostIsAnsweredUnavailable(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	m.Failure = unavailable.Format("mattermost maintenance")
	c, r := servedDaemon(t, m)
	result, e := c.GetChannelPostsWithResponse(
		t.Context(),
		"alfa",
		&client.GetChannelPostsParams{Since: time.Now().Add(-time.Hour)},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusServiceUnavailable, result.StatusCode())
	assert.Count(t, 0, r.Events())
}

func TestAMattermostFaultIsCaptured(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	m.Failure = fault()
	c, r := servedDaemon(t, m)
	result, e := c.GetThreadWithResponse(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, result.StatusCode())
	assert.FatalNotNil(t, result.JSON500)
	assert.Count(t, 1, r.Events())
}
