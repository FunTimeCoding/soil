package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/authorization/authorization_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireHandsSignedInSubjectOn(t *testing.T) {
	c := authorization_tester.Client()
	t.Cleanup(c.Close)
	var seen string
	q := httptest.NewRequest(http.MethodGet, "/target", nil)
	q.AddCookie(c.SubjectCookie("alfa"))
	c.Require(
		func(_ http.ResponseWriter, r *http.Request) { seen = guard.Subject(r) },
	)(httptest.NewRecorder(), q)
	assert.String(t, "alfa", seen)
}
