package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/refusal_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/refusing_source"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"net/http"
	"testing"
)

func TestRejectedCreateIsBadRequestWithoutSentry(t *testing.T) {
	status, events := refusal_tester.Create(t, validation.New("slug: exists"))
	assert.Integer(t, http.StatusBadRequest, status)
	assert.Integer(t, 0, events)
}

func TestRefusedTokenIsCaptured(t *testing.T) {
	status, events := refusal_tester.Create(
		t,
		unexpected.Format("netbox status 403: no permission"),
	)
	assert.Integer(t, http.StatusInternalServerError, status)
	assert.Integer(t, 1, events)
}

func TestConflictWithoutConflictAnswerIsCaptured(t *testing.T) {
	status, events := refusal_tester.Create(
		t,
		conflict.Format("dependent objects"),
	)
	assert.Integer(t, http.StatusInternalServerError, status)
	assert.Integer(t, 1, events)
}

func TestRejectedCreateIsQuietToolFailure(t *testing.T) {
	s := base.New(
		t,
		service_tester.WithClient(
			refusing_source.New(t, validation.New("slug: exists")),
		),
	)
	c := s.Tools()
	assert.StringContains(
		t,
		"slug: exists",
		c.MustCallToolError(
			constant.CreateSite,
			map[string]any{"name": "alfa"},
		),
	)
	assert.Count(t, 0, s.Reporter.Events())
}
