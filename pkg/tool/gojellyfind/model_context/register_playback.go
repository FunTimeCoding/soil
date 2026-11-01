package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerPlayback() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSessions,
			mcp.WithDescription(
				"List active Jellyfin sessions with playback state and device info",
			),
		),
		mcp.NewTypedToolHandler(s.ListSessions),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Play,
			mcp.WithDescription("Instruct a Jellyfin session to play items"),
			mcp.WithString(
				"session_id",
				mcp.Required(),
				mcp.Description("Target session ID"),
			),
			mcp.WithArray(
				"item_ids",
				mcp.WithStringItems(),
				mcp.Required(),
				mcp.Description("Item IDs to play"),
			),
			mcp.WithString(
				"play_command",
				mcp.Description("PlayNow (default), PlayNext, or PlayLast"),
			),
			mcp.WithNumber(
				"start_position_ticks",
				mcp.Description(
					"Starting position in ticks (10,000,000 ticks per second)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.Play),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.PlaybackCommand,
			mcp.WithDescription(
				"Send a playback command to a Jellyfin session: Pause, Unpause, Stop, NextTrack, PreviousTrack, Seek, PlayPause",
			),
			mcp.WithString(
				"session_id",
				mcp.Required(),
				mcp.Description("Target session ID"),
			),
			mcp.WithString(
				"command",
				mcp.Required(),
				mcp.Description(
					"Playback command: Pause, Unpause, Stop, NextTrack, PreviousTrack, Seek, PlayPause",
				),
			),
			mcp.WithNumber(
				"seek_position_ticks",
				mcp.Description(
					"Position to seek to in ticks (required for Seek command)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.PlaybackCommand),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SetVolume,
			mcp.WithDescription(
				"Set the volume level of a Jellyfin session (0-100)",
			),
			mcp.WithString(
				"session_id",
				mcp.Required(),
				mcp.Description("Target session ID"),
			),
			mcp.WithNumber(
				"level",
				mcp.Required(),
				mcp.Description("Volume level (0-100)"),
			),
		),
		mcp.NewTypedToolHandler(s.SetVolume),
	)
}
