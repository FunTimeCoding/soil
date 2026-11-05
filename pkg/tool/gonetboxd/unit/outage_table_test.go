package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"net/http"
	"slices"
	"strings"
	"testing"
)

func TestEveryOperationAnswersANetBoxOutageUnavailable(t *testing.T) {
	r := memory.New()
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	address := `{"address":"00:00:5e:00:53:01"}`
	var unanswered []string

	for operation, status := range outage_tester.Statuses(
		t,
		func() *model_context_server.Server { return outageServer(t, r) },
		spec,
		map[string]string{
			"SetInterfacePhysicalAddress":        address,
			"SetVirtualInterfacePhysicalAddress": address,
		},
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

func TestEveryToolFailsOnANetBoxOutage(t *testing.T) {
	r := memory.New()
	var unanswered []string

	for tool, text := range outage_tester.ToolFailures(
		t,
		func() *model_context_server.Server { return outageServer(t, r) },
		map[string]map[string]any{
			"update_journal_entry": {"comments": "alfa"},
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
