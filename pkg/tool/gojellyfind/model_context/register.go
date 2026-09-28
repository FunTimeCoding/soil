package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListLibraries,
			mcp.WithDescription("List all Jellyfin media libraries"),
		),
		mcp.NewTypedToolHandler(s.ListLibraries),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchItems,
			mcp.WithDescription(
				"Search Jellyfin media items with optional filters",
			),
			mcp.WithString("q", mcp.Description("Free text search")),
			mcp.WithArray(
				"types",
				mcp.WithStringItems(),
				mcp.Description(
					"Item types to include: Movie, Series, Audio, MusicAlbum, MusicArtist",
				),
			),
			mcp.WithNumber("page", mcp.Description("Page number (default: 1)")),
			mcp.WithNumber(
				"per_page",
				mcp.Description("Results per page (default: 25)"),
			),
		),
		mcp.NewTypedToolHandler(s.SearchItems),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetItem,
			mcp.WithDescription("Get a single Jellyfin item by ID"),
			mcp.WithString("id", mcp.Required(), mcp.Description("Item ID")),
		),
		mcp.NewTypedToolHandler(s.GetItem),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetEpisodes,
			mcp.WithDescription("Get all episodes of a Jellyfin series"),
			mcp.WithString(
				"series_id",
				mcp.Required(),
				mcp.Description("Series item ID"),
			),
		),
		mcp.NewTypedToolHandler(s.GetEpisodes),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetTracks,
			mcp.WithDescription("Get all tracks of a Jellyfin music album"),
			mcp.WithString(
				"album_id",
				mcp.Required(),
				mcp.Description("Music album item ID"),
			),
		),
		mcp.NewTypedToolHandler(s.GetTracks),
	)
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
