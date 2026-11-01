package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerMessaging() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Send,
			mcp.WithDescription(
				"Send a message to another session by name, or broadcast to all. Use for coordination: build warnings, scope handoffs, questions.",
			),
			mcp.WithString(
				constant.To,
				mcp.Description("Recipient name (omit to broadcast to all)"),
			),
			mcp.WithString(
				constant.Body,
				mcp.Required(),
				mcp.Description("Message text"),
			),
			mcp.WithString(
				constant.DeliveryParameter,
				mcp.Enum(
					string(constant.DeliveryQueued),
					string(constant.DeliveryImmediate),
					string(constant.DeliveryCruise),
				),
				mcp.Description(
					"queued (default) waits for the recipient's next prompt. immediate wakes them now - only for messages that cannot wait, it costs the recipient a turn the moment it lands. cruise wakes them only as their cruise setting allows.",
				),
			),
		),
		s.send,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ReadMessage,
			mcp.WithDescription(
				"Read coordination messages whole by number. Use it when a delivery cut a message or held it as waiting - the delivery line names the number.",
			),
			mcp.WithArray(
				constant.Identifiers,
				mcp.Required(),
				mcp.WithNumberItems(),
				mcp.Description("Message numbers, one or several"),
			),
		),
		mcp.NewTypedToolHandler(s.readMessage),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Pulse,
			mcp.WithDescription(
				"Send a pulse to the dashboard. Appears on the session card and detail page. Bidirectional - the dashboard can pulse back.",
			),
			mcp.WithString(
				constant.Line,
				mcp.Required(),
				mcp.Description("The line to send"),
			),
		),
		s.pulse,
	)
}
