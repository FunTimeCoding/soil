package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/telemetry/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (b *Builder) Server() *server.MCPServer {
	hooks := &server.Hooks{}

	if b.logger != nil {
		addRequestHooks(hooks, b.logger, b.verbose)
	}

	if b.recorder != nil {
		r := b.recorder
		hooks.AddAfterCallTool(
			func(
				_ context.Context,
				_ any,
				message *mcp.CallToolRequest,
				result any,
			) {
				outcome := constant.Success

				if v, okay := result.(*mcp.CallToolResult); okay && v.IsError {
					outcome = constant.Error
				}

				r.Record(
					record.NewBaseline(
						message.Params.Name,
						constant.ModelContext,
						"model",
						outcome,
					),
				)
			},
		)
	}

	options := []server.ServerOption{
		server.WithToolCapabilities(true),
		server.WithInstructions(b.instructions),
		server.WithHooks(hooks),
		server.WithInputSchemaValidation(),
		server.WithStrictInputSchemaDefault(),
	}

	if b.logger != nil {
		options = append(options, server.WithLogger(b.logger.Slog()))
	}

	if b.resources {
		options = append(
			options,
			server.WithResourceCapabilities(true, false),
		)
	}

	return server.NewMCPServer(b.name, b.version, options...)
}
