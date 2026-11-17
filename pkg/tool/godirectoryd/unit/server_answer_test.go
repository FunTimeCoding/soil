package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/base"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"net/http"
	"testing"
)

func TestOutageIsAnsweredAsUnavailableWithoutCapture(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = outage_tester.RefusedDial()
	c := s.GeneratedClient()

	for range 2 {
		r, e := c.GetUserWithResponse(t.Context())
		assert.FatalOnError(t, e)
		assert.Integer(t, http.StatusServiceUnavailable, r.StatusCode())
		assert.NotNil(t, r.JSON503)
	}

	assert.Count(t, 0, s.Reporter.Events())
}

func TestRefusedServiceAccountIsAnsweredAsForbidden(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = forbidden.Format(
		"directory cn=admin: %s",
		"invalid credentials",
	)
	r, e := s.GeneratedClient().GetGroupWithResponse(t.Context())
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusForbidden, r.StatusCode())
	assert.NotNil(t, r.JSON403)
	assert.Count(t, 0, s.Reporter.Events())
}
