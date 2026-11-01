package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerChannels() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListChannels,
			mcp.WithDescription(
				"List public channels in the Mattermost workspace",
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Max channels to return (default 100)"),
				mcp.DefaultNumber(100),
			),
			mcp.WithNumber(
				"page",
				mcp.Description("Page number starting from 0"),
				mcp.DefaultNumber(0),
			),
		),
		mcp.NewTypedToolHandler(s.ListChannels),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetChannel,
			mcp.WithDescription(
				"Get one channel with its header and purpose. Provide channel_id or channel_name.",
			),
			mcp.WithString("channel_id", mcp.Description("Channel ID")),
			mcp.WithString("channel_name", mcp.Description("Channel name")),
		),
		mcp.NewTypedToolHandler(s.GetChannel),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetChannelHistory,
			mcp.WithDescription(
				"Get recent messages from a channel. Provide channel_id or channel_name.",
			),
			mcp.WithString("channel_id", mcp.Description("Channel ID")),
			mcp.WithString("channel_name", mcp.Description("Channel name")),
			mcp.WithNumber(
				"limit",
				mcp.Description(
					"Number of messages (default 30). Also caps since results, keeping the newest.",
				),
				mcp.DefaultNumber(30),
			),
			mcp.WithString(
				"since",
				mcp.Description(
					"Only messages since this time (e.g. \"2026-04-22 15:00\").",
				),
			),
		),
		mcp.NewTypedToolHandler(s.GetChannelHistory),
	)
}
