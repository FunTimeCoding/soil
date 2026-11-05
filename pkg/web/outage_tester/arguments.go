package outage_tester

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"maps"
)

func arguments(
	tool *mcp.Tool,
	overrides map[string]any,
) map[string]any {
	var s schema
	notation.MustDecodeBytes(notation.Marshal(tool.InputSchema), &s, false)
	result := map[string]any{}

	for _, name := range s.Required {
		result[name] = placeholder(s.Properties[name])
	}

	maps.Copy(result, overrides)

	return result
}
