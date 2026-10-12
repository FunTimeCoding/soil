package anthropic

import (
	"context"
	"fmt"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Run() {
	type (
		Input struct {
			Name string `json:"name" jsonschema:"name of person to greet"`
		}
		Output struct {
			Greeting string `json:"greeting" jsonschema:"greeting to tell the user"`
		}
	)
	s := mcp.NewServer(
		&mcp.Implementation{Name: "Greeter", Version: soil.DefaultVersion},
		nil,
	)
	mcp.AddTool(
		s,
		&mcp.Tool{
			Name:        constant.ModelContextGreetTool,
			Description: "say hi",
		},
		func(
			_ context.Context,
			_ *mcp.CallToolRequest,
			i Input,
		) (
			*mcp.CallToolResult,
			Output,
			error,
		) {
			return nil, Output{Greeting: fmt.Sprintf("Hi %s", i.Name)}, nil
		},
	)
	errors.PanicOnError(s.Run(context.Background(), &mcp.StdioTransport{}))
}
