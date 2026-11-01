package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerLibrary() {
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
}
