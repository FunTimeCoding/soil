package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"testing"
)

func recoveringClient(
	t *testing.T,
	r face.Reporter,
) *client.Client {
	t.Helper()
	s := server.New(identity.Example(), r).WithResources().Server()
	s.AddTool(
		mcp.NewTool("golf"),
		func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			panic("bravo")
		},
	)
	s.AddResource(
		mcp.NewResource("test://charlie", "charlie"),
		func(
			context.Context,
			mcp.ReadResourceRequest,
		) ([]mcp.ResourceContents, error) {
			panic("delta")
		},
	)
	result, e := client.NewInProcessClient(s)
	assert.FatalOnError(t, e)
	t.Cleanup(func() { assert.FatalOnError(t, result.Close()) })
	assert.FatalOnError(t, result.Start(t.Context()))
	_, f := result.Initialize(t.Context(), mcp.InitializeRequest{})
	assert.FatalOnError(t, f)

	return result
}
