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
	s.server.AddTool(
		mcp.NewTool(
			constant.SendChatMessage,
			mcp.WithDescription(
				"Send a message to a chat as the signed-in user; line breaks are kept",
			),
			mcp.WithString(
				"chat",
				mcp.Required(),
				mcp.Description("Chat identifier from my_chats"),
			),
			mcp.WithString(
				"text",
				mcp.Required(),
				mcp.Description(
					"Message text, plain; write @Name where a mentioned member should be tagged",
				),
			),
			mcp.WithArray(
				"mention",
				mcp.Description(
					"Display names of chat members to tag; each must appear as @Name in the text",
				),
			),
		),
		mcp.NewTypedToolHandler(s.SendChatMessage),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SendDirectMessage,
			mcp.WithDescription(
				"Send a direct message to a person, opening the one-on-one chat if there is none yet",
			),
			mcp.WithString(
				"user",
				mcp.Required(),
				mcp.Description("Display name, email address or identifier"),
			),
			mcp.WithString(
				"text",
				mcp.Required(),
				mcp.Description(
					"Message text, plain; write @Name where a mentioned member should be tagged",
				),
			),
			mcp.WithArray(
				"mention",
				mcp.Description(
					"Display names to tag; each must appear as @Name in the text",
				),
			),
		),
		mcp.NewTypedToolHandler(s.SendDirectMessage),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateChat,
			mcp.WithDescription(
				"Open a chat with the signed-in user and the given people: one person makes a one-on-one chat, more make a group chat",
			),
			mcp.WithArray(
				"members",
				mcp.Required(),
				mcp.Description("Display names, email addresses or identifiers"),
			),
			mcp.WithString("topic", mcp.Description("Topic of a group chat")),
		),
		mcp.NewTypedToolHandler(s.CreateChat),
	)
}
