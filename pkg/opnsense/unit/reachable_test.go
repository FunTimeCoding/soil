package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/opnsense/unit/opnsense_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestProbeAnswersWhenOpnsenseAnswers(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithTLS(),
		upstream_server.WithAnswer("{}"),
	)
	assert.FatalOnError(t, opnsense_tester.New(t, s.URL).Reachable(t.Context()))
}

func TestDeadKeyFailsProbe(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithTLS(),
		upstream_server.WithStatus(http.StatusUnauthorized),
	)
	assert.NotNil(t, opnsense_tester.New(t, s.URL).Reachable(t.Context()))
}
