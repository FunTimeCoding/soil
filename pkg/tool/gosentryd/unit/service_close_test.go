package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/unit/service_tester"
	"testing"
)

func TestCloseClosesTheSentryClient(t *testing.T) {
	c := service_tester.NewSource(t)
	service_tester.Service(t, c).Close()
	assert.Integer(t, 1, c.Closed)
}

func TestWaitAnswersOnceSentryAnswers(t *testing.T) {
	assert.FatalOnError(
		t,
		service_tester.Service(t, service_tester.NewSource(t)).Wait(t.Context()),
	)
}
