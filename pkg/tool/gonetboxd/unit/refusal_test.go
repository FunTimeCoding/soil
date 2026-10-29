package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/refusing_source"
	"net/http"
	"testing"
)

func TestARejectedCreateIsBadRequestWithoutSentry(t *testing.T) {
	status, events := refusedCreate(t, validation.New("slug: exists"))
	assert.Integer(t, http.StatusBadRequest, status)
	assert.Integer(t, 0, events)
}

func TestARefusedTokenIsForbiddenWithoutSentry(t *testing.T) {
	status, events := refusedCreate(t, forbidden.Format("no permission"))
	assert.Integer(t, http.StatusForbidden, status)
	assert.Integer(t, 0, events)
}

func TestAConflictWithoutAConflictAnswerIsCaptured(t *testing.T) {
	status, events := refusedCreate(t, conflict.Format("dependent objects"))
	assert.Integer(t, http.StatusInternalServerError, status)
	assert.Integer(t, 1, events)
}

func TestARejectedCreateIsAQuietToolFailure(t *testing.T) {
	r := memory.New()
	s := newNetboxServer(
		t,
		refusing_source.New(validation.New("slug: exists")),
		r,
	)
	c := s.Connect(t)
	defer c.Close()
	assert.StringContains(
		t,
		"slug: exists",
		c.MustCallToolError(constant.CreateSite, map[string]any{"name": "alfa"}),
	)
	assert.Count(t, 0, r.Events())
}
