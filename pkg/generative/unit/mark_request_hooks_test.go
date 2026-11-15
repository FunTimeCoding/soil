package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/unit/request_tester"
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
		logged := request_tester.Logged(m, message)
		assert.False(t, strings.Contains(logged, "level=ERROR"))
		assert.StringContains(t, "level=INFO", logged)
		assert.StringContains(t, string(m), logged)
	}
}

func TestToolCallIsLoggedWithItsName(t *testing.T) {
	q := &mcp.CallToolRequest{}
	q.Method = string(mcp.MethodToolsCall)
	q.Params.Name = "roll"
	assert.StringContains(
		t,
		"tools/call roll",
		request_tester.Logged(mcp.MethodToolsCall, q),
	)
}
