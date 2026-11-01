package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerJournalEdits() {
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateJournalEntry,
			mcp.WithDescription("Update a journal entry's comments or kind"),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Journal entry identifier"),
			),
			mcp.WithString(
				constant.Comments,
				mcp.Description("New entry text, markdown"),
			),
			mcp.WithString(
				constant.Kind,
				mcp.Description(
					"New entry kind: info, success, warning, or danger",
				),
			),
		),
		s.updateJournalEntry,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteJournalEntry,
			mcp.WithDescription("Delete a journal entry"),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Journal entry identifier"),
			),
		),
		s.deleteJournalEntry,
	)
}
