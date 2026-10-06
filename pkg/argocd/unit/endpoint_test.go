package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"testing"
)

func TestApplicationsSendTheToken(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"items":[]}`)
	_, e := newClient(t, s.URL).Applications()
	assert.FatalOnError(t, e)
	r := last()
	assert.String(t, "/api/v1/applications", r.URL.Path)
	assert.String(t, "Bearer alfa-token", r.Header.Get(constant.Authorization))
}

func TestRefusalCarriesTheReason(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusForbidden,
		`{"error":"permission denied","code":7,"message":"permission denied"}`,
	)
	_, e := newClient(t, s.URL).Applications()
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "status: 403: permission denied", e.Error())
}
