package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerConfluenceEdits() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceEditPage,
			mcp.WithDescription(
				"Edit a Confluence page by replacing a text fragment. Call confluence_get_page first to read the current content. The old_text must appear exactly once in the page.",
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithString(
				constant.OldText,
				mcp.Required(),
				mcp.Description(
					"The exact text to find and replace. Must match exactly once.",
				),
			),
			mcp.WithString(
				constant.NewText,
				mcp.Required(),
				mcp.Description("The replacement text"),
			),
			mcp.WithString(
				generativeConstant.ParameterTitle,
				mcp.Description(
					"New page title. If omitted, the current title is kept.",
				),
			),
			mcp.WithString(
				generativeConstant.ParameterMessage,
				mcp.Description("Version comment"),
			),
			mcp.WithBoolean(
				constant.Draft,
				mcp.Description(
					"Edit a draft page. The page stays as a draft after editing.",
				),
			),
		),
		s.editPage,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceAddComment,
			mcp.WithDescription("Add a comment to a Confluence page"),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithString(
				generativeConstant.ParameterBody,
				mcp.Required(),
				mcp.Description("Comment text"),
			),
		),
		s.addPageComment,
	)
}
