package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerConfluencePages() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceCreatePage,
			mcp.WithDescription(
				"Create a new Confluence page with markdown content",
			),
			mcp.WithString(
				constant.SpaceIdentifier,
				mcp.Required(),
				mcp.Description("Space ID"),
			),
			mcp.WithString(
				constant.ParentIdentifier,
				mcp.Required(),
				mcp.Description("Parent page ID"),
			),
			mcp.WithString(
				generativeConstant.ParameterTitle,
				mcp.Description(
					"Page title. Required for published pages, optional for drafts.",
				),
			),
			mcp.WithString(
				generativeConstant.ParameterBody,
				mcp.Required(),
				mcp.Description("Page content in markdown"),
			),
			mcp.WithBoolean(
				constant.Draft,
				mcp.Description(
					"Create as a draft page instead of publishing immediately.",
				),
			),
		),
		s.createPage,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceUpdatePage,
			mcp.WithDescription(
				"Update a Confluence page with markdown content. Gets the current version automatically.",
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithString(
				generativeConstant.ParameterTitle,
				mcp.Required(),
				mcp.Description("Page title"),
			),
			mcp.WithString(
				generativeConstant.ParameterBody,
				mcp.Required(),
				mcp.Description("Full page content in markdown"),
			),
			mcp.WithString(
				generativeConstant.ParameterMessage,
				mcp.Description("Version comment"),
			),
		),
		s.updatePage,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceDeletePage,
			mcp.WithDescription(
				"Delete a Confluence page. For draft pages, set draft to true - draft pages are permanently deleted, not sent to trash. Draft overlays (unpublished changes on published pages) cannot be deleted via the API.",
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithBoolean(
				constant.Draft,
				mcp.Description(
					"Delete a draft page. Required for pages that have never been published.",
				),
			),
		),
		s.deletePage,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceSetPageStatus,
			mcp.WithDescription(
				"Publish or unpublish a Confluence page. Set status to 'current' to publish a draft, or 'draft' to unpublish a published page.",
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithString(
				constant.PageStatus,
				mcp.Required(),
				mcp.Description(
					"Target status: 'current' to publish, 'draft' to unpublish.",
				),
			),
		),
		s.setPageStatus,
	)
}
