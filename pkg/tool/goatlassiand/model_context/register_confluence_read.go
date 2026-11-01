package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerConfluenceRead() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceSearch,
			mcp.WithDescription(
				"Search Confluence pages using CQL or plain text",
			),
			mcp.WithString(
				generativeConstant.ParameterQuery,
				mcp.Required(),
				mcp.Description("CQL query string or plain text"),
			),
		),
		s.searchPages,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceGetPage,
			mcp.WithDescription(
				"Get a Confluence page by ID with body content as markdown",
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithBoolean(
				constant.Draft,
				mcp.Description(
					"Fetch a draft page. Required for unpublished pages which return 'not found' without this flag.",
				),
			),
		),
		s.getPage,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceGetPageDraft,
			mcp.WithDescription(
				"Get the draft layer of a Confluence page. Confluence maintains a draft layer for every page - this always returns content, even if there are no unpublished changes. Compare against the published version (confluence_get_page) to detect actual differences.",
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
		),
		s.getPageDraft,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceListPages,
			mcp.WithDescription(
				"List pages in a Confluence space. Defaults to published pages. Set status to 'draft' to find draft pages - returns all drafts visible to the current user (the v1 API does not filter drafts by space).",
			),
			mcp.WithString(
				constant.SpaceIdentifier,
				mcp.Required(),
				mcp.Description("Space ID"),
			),
			mcp.WithString(
				constant.PageStatus,
				mcp.Description(
					"Filter by status: 'current' (default) for published pages, 'draft' for draft pages.",
				),
			),
		),
		s.listPages,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceListSpaces,
			mcp.WithDescription("List all visible Confluence spaces"),
		),
		s.listSpaces,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceGetPageChildren,
			mcp.WithDescription("List child pages of a Confluence page"),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Parent page ID"),
			),
		),
		s.getPageChildren,
	)
}
