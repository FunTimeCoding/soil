package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"testing"
)

func SendToolNotation(
	t *testing.T,
	tools []*mcp.Tool,
) string {
	t.Helper()

	for _, tool := range tools {
		if tool.Name == constant.Send {
			return string(notation.Marshal(tool))
		}
	}

	t.Fatalf("no %s tool listed", constant.Send)

	return ""
}
