package session_tester

import "github.com/mark3labs/mcp-go/mcp"

type Session struct {
	identifier    string
	notifications chan mcp.JSONRPCNotification
}
