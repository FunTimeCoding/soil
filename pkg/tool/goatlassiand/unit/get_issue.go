package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generated "github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/failing_jira"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	"testing"
)

func getIssue(
	t *testing.T,
	failure error,
) (generated.GetIssueResponseObject, *memory.Memory) {
	t.Helper()
	m := memory.New()
	r, e := server.New(
		service_tester.NewService(
			t,
			failing_jira.New(failure),
			service_tester.NewConfluence(t),
			nil,
		),
		m,
	).GetIssue(
		context.Background(),
		generated.GetIssueRequestObject{Key: "ALFA-1"},
	)
	assert.FatalOnError(t, e)

	return r, m
}
