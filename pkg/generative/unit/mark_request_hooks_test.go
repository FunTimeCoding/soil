package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
	"testing"
)

func TestEveryRequestIsLoggedByItsMethod(t *testing.T) {
	for m, message := range map[mcp.MCPMethod]any{
		mcp.MethodResourcesList:          &mcp.ListResourcesRequest{},
		mcp.MethodResourcesTemplatesList: &mcp.ListResourceTemplatesRequest{},
		mcp.MethodResourcesRead:          &mcp.ReadResourceRequest{},
		mcp.MethodPromptsList:            &mcp.ListPromptsRequest{},
		mcp.MethodPing:                   &mcp.PingRequest{},
		mcp.MethodInitialize:             &mcp.InitializeRequest{},
		mcp.MethodToolsList:              &mcp.ListToolsRequest{},
	} {
		logged := hookLog(m, message)
		assert.False(t, strings.Contains(logged, "level=ERROR"))
		assert.StringContains(t, "level=INFO", logged)
		assert.StringContains(t, string(m), logged)
	}
}

func TestToolCallIsLoggedWithItsName(t *testing.T) {
	q := &mcp.CallToolRequest{}
	q.Method = string(mcp.MethodToolsCall)
	q.Params.Name = "roll"
	assert.StringContains(t, "tools/call roll", hookLog(mcp.MethodToolsCall, q))
}
