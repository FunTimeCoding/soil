package session_tester

import "github.com/mark3labs/mcp-go/mcp"

func New(identifier string) *Session {
	return &Session{
		identifier:    identifier,
		notifications: make(chan mcp.JSONRPCNotification, 1),
	}
}
