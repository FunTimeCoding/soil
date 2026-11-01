package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerMessages() {
	s.server.AddTool(
		mcp.NewTool(
			constant.SendDirectMessage,
			mcp.WithDescription("Send a direct message to a user by username."),
			mcp.WithString(
				"username",
				mcp.Required(),
				mcp.Description("Recipient's username"),
			),
			mcp.WithString(
				"message",
				mcp.Required(),
				mcp.Description("Message text"),
			),
		),
		mcp.NewTypedToolHandler(s.SendDirectMessage),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.PostMessage,
			mcp.WithDescription("Post a message to a Mattermost channel"),
			mcp.WithString("channel_id", mcp.Description("Channel ID")),
			mcp.WithString("channel_name", mcp.Description("Channel name")),
			mcp.WithString(
				"message",
				mcp.Required(),
				mcp.Description("Message text"),
			),
		),
		mcp.NewTypedToolHandler(s.PostMessage),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ReplyToThread,
			mcp.WithDescription(
				"Reply to a message thread, optionally adding an emoji reaction to the root message",
			),
			mcp.WithString(
				"channel_id",
				mcp.Required(),
				mcp.Description("Channel ID"),
			),
			mcp.WithString(
				"post_id",
				mcp.Required(),
				mcp.Description("Parent post ID"),
			),
			mcp.WithString(
				"message",
				mcp.Required(),
				mcp.Description("Reply text"),
			),
			mcp.WithString(
				"emoji_name",
				mcp.Description(
					"Emoji reaction without colons to add to the root message",
				),
			),
		),
		mcp.NewTypedToolHandler(s.ReplyToThread),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AddReaction,
			mcp.WithDescription("Add an emoji reaction to a post"),
			mcp.WithString(
				"post_id",
				mcp.Required(),
				mcp.Description("Post ID"),
			),
			mcp.WithString(
				"emoji_name",
				mcp.Required(),
				mcp.Description("Emoji name without colons"),
			),
		),
		mcp.NewTypedToolHandler(s.AddReaction),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetThreadReplies,
			mcp.WithDescription("Get all replies in a message thread"),
			mcp.WithString(
				"post_id",
				mcp.Required(),
				mcp.Description("Parent post ID"),
			),
		),
		mcp.NewTypedToolHandler(s.GetThreadReplies),
	)
}
