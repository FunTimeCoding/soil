package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"net/http"
	"testing"
)

func TestProbeAnswersWhenOpnsenseAnswers(t *testing.T) {
	s := canned("{}")
	defer s.Close()
	assert.FatalOnError(t, client(t, s).Reachable(t.Context()))
}

func TestDeadKeyFailsProbe(t *testing.T) {
	s := refusing(http.StatusUnauthorized)
	defer s.Close()
	assert.NotNil(t, client(t, s).Reachable(t.Context()))
}
