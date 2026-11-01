package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerDecisions() {
	s.server.AddTool(
		mcp.NewTool(
			constant.PushDecision,
			mcp.WithDescription(
				"Hand one decision to the queue without spending a conversation turn on it. Use this for every decision past the first in a message - a reply runs about twenty words however much is put in front of the person answering, so anything beyond one ask riding along in prose gets a 'yes' that is not an approval. The question must fit in 300 characters and a default action is required.",
			),
			mcp.WithString(
				constant.SessionParameter,
				mcp.Required(),
				mcp.Description("Your coordination callsign"),
			),
			mcp.WithString(
				constant.QuestionParameter,
				mcp.Required(),
				mcp.Description("The decision, in 300 characters or fewer"),
			),
			mcp.WithString(
				constant.DefaultActionParameter,
				mcp.Required(),
				mcp.Description(
					"What you will do if no answer comes. If this is easy to write and obviously fine, do not push the decision at all - just do it.",
				),
			),
			mcp.WithArray(
				constant.OptionParameter,
				mcp.Description(
					"Short answer options, shorter than an interview's",
				),
			),
			mcp.WithArray(
				constant.FrameParameter,
				mcp.Description(
					"Topics this decision belongs to; a decision may sit in several",
				),
			),
		),
		mcp.NewTypedToolHandler(s.pushDecision),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DecisionStatus,
			mcp.WithDescription(
				"Poll what the queue has handled. Returns open counts per frame and the decisions answered, declined or called irrelevant that you have not cleared yet.",
			),
			mcp.WithString(
				constant.SessionParameter,
				mcp.Required(),
				mcp.Description("Your coordination callsign"),
			),
		),
		mcp.NewTypedToolHandler(s.decisionStatus),
	)
}
