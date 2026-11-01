package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerNotes() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListNotes,
			mcp.WithDescription("List notes on a link."),
			mcp.WithNumber(
				"link_identifier",
				mcp.Required(),
				mcp.Description("Link ID to list notes for."),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of results to return."),
			),
			mcp.WithNumber(
				"offset",
				mcp.Description("Number of results to skip."),
			),
		),
		mcp.NewTypedToolHandler(s.ListNotes),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AddNote,
			mcp.WithDescription("Add a note to a link."),
			mcp.WithNumber(
				"link_identifier",
				mcp.Required(),
				mcp.Description("Link ID to add the note to."),
			),
			mcp.WithString(
				"text",
				mcp.Required(),
				mcp.Description("Note text."),
			),
		),
		mcp.NewTypedToolHandler(s.AddNote),
	)
}
