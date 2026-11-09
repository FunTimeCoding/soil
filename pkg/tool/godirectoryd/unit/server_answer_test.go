package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/service_tester"
	"net/http"
	"testing"
)

func TestOutageIsAnsweredAsUnavailableWithoutCapture(t *testing.T) {
	c := service_tester.NewSource()
	c.Failure = refusedConnection()
	m := memory.New()
	s := restClient(t, c, m)

	for range 2 {
		r, e := s.GetUserWithResponse(t.Context())
		assert.FatalOnError(t, e)
		assert.Integer(t, http.StatusServiceUnavailable, r.StatusCode())
		assert.NotNil(t, r.JSON503)
	}

	assert.Count(t, 0, m.Events())
}

func TestRefusedServiceAccountIsAnsweredAsForbidden(t *testing.T) {
	c := service_tester.NewSource()
	c.Failure = forbidden.Format(
		"directory cn=admin: %s",
		"invalid credentials",
	)
	m := memory.New()
	r, e := restClient(t, c, m).GetGroupWithResponse(t.Context())
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusForbidden, r.StatusCode())
	assert.NotNil(t, r.JSON403)
	assert.Count(t, 0, m.Events())
}
