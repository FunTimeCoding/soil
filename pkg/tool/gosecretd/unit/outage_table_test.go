package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/server"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"net/http"
	"slices"
	"strings"
	"testing"
)

func TestEveryOperationAnswersOpenbaoOutageUnavailable(t *testing.T) {
	r := memory.New()
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	var unanswered []string

	for operation, status := range outage_tester.Statuses(
		t,
		func() *model_context_server.Server { return outageServer(t, r) },
		spec,
		map[string]string{"PutSecret": `{"values":{"alfa":"bravo"}}`},
	) {
		if status != http.StatusServiceUnavailable {
			unanswered = append(
				unanswered,
				fmt.Sprintf("%s %d", operation, status),
			)
		}
	}

	slices.Sort(unanswered)
	assert.String(t, "", strings.Join(unanswered, ", "))
	assert.Count(t, 0, r.Events())
}

func TestEveryToolFailsOnOpenbaoOutage(t *testing.T) {
	r := memory.New()
	var unanswered []string

	for tool, text := range outage_tester.ToolFailures(
		t,
		func() *model_context_server.Server { return outageServer(t, r) },
		map[string]map[string]any{
			"put_secret": {"values": map[string]any{"alfa": "bravo"}},
		},
	) {
		if !strings.Contains(text, "connection refused") {
			unanswered = append(unanswered, tool)
		}
	}
	slices.Sort(unanswered)
	assert.String(t, "", strings.Join(unanswered, ", "))
	assert.Count(t, 0, r.Events())
}
