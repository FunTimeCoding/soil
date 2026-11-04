package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/sentry/trace_tester"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"testing"
)

func TestAnswerWithCarriesTheTraceOfItsContext(t *testing.T) {
	m := memory.New()
	_, e := response.AnswerWith(
		trace_tester.Context(t, nil),
		m,
		errors.New("disk full"),
		"load post",
	)
	assert.Nil(t, e)
	assert.Count(t, 1, m.Events())
	assert.StringContains(
		t,
		"0123456789abcdef0123456789abcdef",
		m.Events()[0].Trace,
	)
}

func TestCaptureFailCarriesTheTraceOfItsContext(t *testing.T) {
	m := memory.New()
	_, e := response.CaptureFail(
		trace_tester.Context(t, nil),
		m,
		errors.New("disk full"),
		"load post",
	)
	assert.Nil(t, e)
	assert.Count(t, 1, m.Events())
	assert.StringContains(
		t,
		"0123456789abcdef0123456789abcdef",
		m.Events()[0].Trace,
	)
}
