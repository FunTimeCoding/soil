package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerDecisionFollowUps() {
	s.server.AddTool(
		mcp.NewTool(
			constant.AddTurn,
			mcp.WithDescription(
				"Follow up on one pushed decision. Capped at 300 characters and three turns - past that, clear it or push a new decision under a new framing.",
			),
			mcp.WithNumber(
				constant.IdentifierParameter,
				mcp.Required(),
				mcp.Description("Decision identifier"),
			),
			mcp.WithString(
				constant.ContentParameter,
				mcp.Required(),
				mcp.Description("The follow-up, in 300 characters or fewer"),
			),
		),
		mcp.NewTypedToolHandler(s.addTurn),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ClearDecision,
			mcp.WithDescription(
				"Close a decision by saying in one line what you understood. Whether it resolved on an answer or on your default is recorded from the stored state, not from what you claim.",
			),
			mcp.WithNumber(
				constant.IdentifierParameter,
				mcp.Required(),
				mcp.Description("Decision identifier"),
			),
			mcp.WithString(
				constant.UnderstandingParameter,
				mcp.Required(),
				mcp.Description("One line: what you took from it"),
			),
			mcp.WithString(
				constant.HeardParameter,
				mcp.Description(
					"Only when the answer came in conversation instead of through the queue: your paraphrase of it. Recorded as the answer, marked as arriving through conversation rather than a click. Ignored when the queue already holds an answer - a click is never overwritten by your transcription.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.clearDecision),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.BumpDecision,
			mcp.WithDescription(
				"Mark one decision as the one blocking you most. This is a counter shown on the queue, not a notification - it says which decision is holding you up, which nothing else shows.",
			),
			mcp.WithNumber(
				constant.IdentifierParameter,
				mcp.Required(),
				mcp.Description("Decision identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.bumpDecision),
	)
}
