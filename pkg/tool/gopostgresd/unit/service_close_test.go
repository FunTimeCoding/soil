package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/unit/service_tester"
	"testing"
)

func TestClosingTheServiceClosesEveryInstanceClient(t *testing.T) {
	s, clients := service_tester.New(t, "alfa", "bravo")
	s.Close()
	assert.Integer(t, 1, clients["alfa"].Closed)
	assert.Integer(t, 1, clients["bravo"].Closed)
}

func TestTheServiceHandsOneGuardPerInstanceToTheLifecycle(t *testing.T) {
	s, _ := service_tester.New(t, "alfa", "bravo")
	assert.Count(t, 2, s.Workers())
}
