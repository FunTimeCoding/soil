package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/service_tester"
	"testing"
)

func TestCloseClosesTheLinkAceClient(t *testing.T) {
	c := service_tester.NewSource(t)
	service_tester.Service(t, c, memory.New()).Close()
	assert.Integer(t, 1, c.Closed)
}

func TestWaitAnswersOnceLinkAceAnswers(t *testing.T) {
	assert.FatalOnError(
		t,
		service_tester.Service(t, service_tester.NewSource(t), memory.New()).
			Wait(t.Context()),
	)
}
