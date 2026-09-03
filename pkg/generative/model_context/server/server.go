package server

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/mark3labs/mcp-go/server"
	"log/slog"
	"sync"
)

type Server struct {
	server               *server.MCPServer
	logger               *slog.Logger
	context              context.Context
	openAuthentication   bool
	verifier             *oidc.IDTokenVerifier
	once                 sync.Once
	serverLocator        string
	authorizationLocator string
	clientIdentifier     string
}
