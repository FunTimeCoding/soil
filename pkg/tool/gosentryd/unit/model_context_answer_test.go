package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/unit/service_tester"
	"github.com/mark3labs/mcp-go/mcp"
	"net"
	"os"
	"syscall"
	"testing"
)

func TestAnUnknownIssueIsAnsweredQuietly(t *testing.T) {
	m := memory.New()
	s := model_context.New(
		service_tester.Service(t, service_tester.NewSource(t)),
		m,
		mock_recorder.New(),
	)
	r, e := s.GetIssue(
		context.Background(),
		mcp.CallToolRequest{},
		argument.GetIssue{Identifier: "100"},
	)
	assert.FatalOnError(t, e)
	assert.True(t, r.IsError)
	assert.String(
		t,
		"issue not found: 100",
		r.Content[0].(mcp.TextContent).Text,
	)
	assert.Count(t, 0, m.Events())
}

func TestAForbiddenSearchIsAnsweredQuietly(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = forbidden.Format("sentry: %s", "token lacks event:read")
	m := memory.New()
	r, e := model_context.New(
		service_tester.Service(t, c),
		m,
		mock_recorder.New(),
	).SearchIssues(
		context.Background(),
		mcp.CallToolRequest{},
		argument.SearchIssues{},
	)
	assert.FatalOnError(t, e)
	assert.True(t, r.IsError)
	assert.Count(t, 0, m.Events())
}

func TestAnUnclassifiedFailureIsCaptured(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = errors.New("sentry answered nonsense")
	m := memory.New()
	r, e := model_context.New(
		service_tester.Service(t, c),
		m,
		mock_recorder.New(),
	).Whoami(
		context.Background(),
		mcp.CallToolRequest{},
		argument.Whoami{},
	)
	assert.FatalOnError(t, e)
	assert.True(t, r.IsError)
	assert.Count(t, 1, m.Events())
}

func TestAnOutageIsAnsweredWithoutACapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
	}
	m := memory.New()
	s := model_context.New(service_tester.Service(t, c), m, mock_recorder.New())

	for range 2 {
		r, e := s.Whoami(
			context.Background(),
			mcp.CallToolRequest{},
			argument.Whoami{},
		)
		assert.FatalOnError(t, e)
		assert.True(t, r.IsError)
	}

	assert.Count(t, 0, m.Events())
}
