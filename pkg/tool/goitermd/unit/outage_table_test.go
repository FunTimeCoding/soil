package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"slices"
	"strings"
	"testing"
)

func TestEveryToolFailsOnItermOutage(t *testing.T) {
	r := memory.New()
	var unanswered []string

	for tool, text := range outage_tester.ToolFailures(
		t,
		func() *model_context_server.Server { return outageServer(t, r) },
		nil,
	) {
		if !strings.Contains(text, "connection refused") {
			unanswered = append(unanswered, tool)
		}
	}

	slices.Sort(unanswered)
	assert.String(t, "", strings.Join(unanswered, ", "))
	assert.Count(t, 0, r.Events())
}
