package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gographd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerTeams() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTeams,
			mcp.WithDescription(
				"List the Teams teams the signed-in user belongs to",
			),
		),
		mcp.NewTypedToolHandler(s.ListTeams),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListChannels,
			mcp.WithDescription("List the channels of a team"),
			mcp.WithString(
				"team",
				mcp.Required(),
				mcp.Description("Team name or identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.ListChannels),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetChannelHistory,
			mcp.WithDescription(
				"Read a channel's threads, newest activity first: each root message with author, time, text and reply count",
			),
			mcp.WithString(
				"team",
				mcp.Required(),
				mcp.Description("Team name or identifier"),
			),
			mcp.WithString(
				"channel",
				mcp.Required(),
				mcp.Description("Channel name or identifier"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of threads (default 20)"),
			),
		),
		mcp.NewTypedToolHandler(s.GetChannelHistory),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetThreadReplies,
			mcp.WithDescription("Read the replies to a channel message"),
			mcp.WithString(
				"team",
				mcp.Required(),
				mcp.Description("Team name or identifier"),
			),
			mcp.WithString(
				"channel",
				mcp.Required(),
				mcp.Description("Channel name or identifier"),
			),
			mcp.WithString(
				"message",
				mcp.Required(),
				mcp.Description("Identifier of the root message"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of replies (default 20)"),
			),
		),
		mcp.NewTypedToolHandler(s.GetThreadReplies),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.PostMessage,
			mcp.WithDescription(
				"Post a new top-level message to a channel as the signed-in user; line breaks are kept",
			),
			mcp.WithString(
				"team",
				mcp.Required(),
				mcp.Description("Team name or identifier"),
			),
			mcp.WithString(
				"channel",
				mcp.Required(),
				mcp.Description("Channel name or identifier"),
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
					"Display names of channel members to tag; each must appear as @Name in the text",
				),
			),
		),
		mcp.NewTypedToolHandler(s.PostMessage),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ReplyToThread,
			mcp.WithDescription(
				"Reply to a channel message as the signed-in user; line breaks are kept",
			),
			mcp.WithString(
				"team",
				mcp.Required(),
				mcp.Description("Team name or identifier"),
			),
			mcp.WithString(
				"channel",
				mcp.Required(),
				mcp.Description("Channel name or identifier"),
			),
			mcp.WithString(
				"message",
				mcp.Required(),
				mcp.Description("Identifier of the root message"),
			),
			mcp.WithString(
				"text",
				mcp.Required(),
				mcp.Description(
					"Reply text, plain; write @Name where a mentioned member should be tagged",
				),
			),
			mcp.WithArray(
				"mention",
				mcp.Description(
					"Display names of channel members to tag; each must appear as @Name in the text",
				),
			),
		),
		mcp.NewTypedToolHandler(s.ReplyToThread),
	)
}
