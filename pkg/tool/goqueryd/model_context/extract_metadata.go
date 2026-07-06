package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func extractMetadata(q mcp.CallToolRequest) map[string]string {
	arguments := q.GetArguments()
	raw, okay := arguments[constant.Metadata]

	if !okay {
		return nil
	}

	m, okay := raw.(map[string]any)

	if !okay {
		return nil
	}

	result := map[string]string{}

	for key, value := range m {
		s, okay := value.(string)

		if okay {
			result[key] = s
		}
	}

	if len(result) == 0 {
		return nil
	}

	return result
}
