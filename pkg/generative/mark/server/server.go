package server

import (
	"context"
	telemetryConstant "github.com/funtimecoding/soil/pkg/telemetry/constant"
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
				outcome := telemetryConstant.Success

				if v, okay := result.(*mcp.CallToolResult); okay && v.IsError {
					outcome = telemetryConstant.Error
				}

				r.Record(
					record.NewBaseline(
						message.Params.Name,
						telemetryConstant.ModelContext,
						"model",
						outcome,
					),
				)
			},
		)
	}

	if b.connected != nil {
		connected := b.connected
		hooks.AddAfterInitialize(
			func(
				_ context.Context,
				_ any,
				_ *mcp.InitializeRequest,
				_ *mcp.InitializeResult,
			) {
				connected()
			},
		)
	}

	if b.sessionEnd != nil {
		sessionEnd := b.sessionEnd
		hooks.AddOnUnregisterSession(
			func(
				_ context.Context,
				s server.ClientSession,
			) {
				sessionEnd(s.SessionID())
			},
		)
	}

	options := []server.ServerOption{
		server.WithToolCapabilities(true),
		server.WithInstructions(b.instructions),
		server.WithHooks(hooks),
		server.WithInputSchemaValidation(),
		server.WithStrictInputSchemaDefault(),
		server.WithToolHandlerMiddleware(recovered(b.reporter)),
		server.WithResourceHandlerMiddleware(recoveredResource(b.reporter)),
	}

	if b.logger != nil {
		options = append(options, server.WithLogger(b.logger.Slog()))
	}

	if b.resources {
		options = append(options, server.WithResourceCapabilities(true, false))
	}

	if b.experimental != nil {
		options = append(options, server.WithExperimental(b.experimental))
	}

	return server.NewMCPServer(b.name, b.version, options...)
}
