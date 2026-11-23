package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gographd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerChats() {
	s.server.AddTool(
		mcp.NewTool(
			constant.MyChats,
			mcp.WithDescription(
				"List the signed-in user's one-to-one, group and meeting chats, most recent message first",
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of chats (default 20)"),
			),
		),
		mcp.NewTypedToolHandler(s.MyChats),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetChatHistory,
			mcp.WithDescription("Read a chat's messages, newest first"),
			mcp.WithString(
				"chat",
				mcp.Required(),
				mcp.Description("Chat identifier from my_chats"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of messages (default 20)"),
			),
		),
		mcp.NewTypedToolHandler(s.GetChatHistory),
	)
}
