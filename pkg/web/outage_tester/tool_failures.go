package outage_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"testing"
)

func ToolFailures(
	t *testing.T,
	serve func() *model_context_server.Server,
	overrides map[string]map[string]any,
) map[string]string {
	t.Helper()
	result := map[string]string{}
	listed := serve()
	listing := listed.Connect(t)
	tools := listing.ListTools()
	listing.Close()
	listed.Close()

	for _, tool := range tools {
		s := serve()
		c := s.Connect(t)
		text, e := c.CallToolError(
			tool.Name,
			arguments(tool, overrides[tool.Name]),
		)
		c.Close()
		s.Close()

		if e != nil {
			text = ""
		}

		result[tool.Name] = text
	}

	if len(result) == 0 {
		t.Fatal("the server lists no tools")
	}

	return result
}
