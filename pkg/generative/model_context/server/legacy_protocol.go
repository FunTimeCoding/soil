package server

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func LegacyProtocol(x context.Context) context.Context {
	return server.WithSupportedProtocolVersions(x, mcp.LegacyProtocolVersions())
}
