package session_tester

import "github.com/mark3labs/mcp-go/mcp"

func (s *Session) NotificationChannel() chan<- mcp.JSONRPCNotification {
	return s.notifications
}
