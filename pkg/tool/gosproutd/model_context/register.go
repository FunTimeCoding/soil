package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSeeds,
			mcp.WithDescription(
				"List all seeds sorted by priority. Each seed carries its file modification time.",
			),
		),
		mcp.NewTypedToolHandler(s.listSeeds),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ReorderSeed,
			mcp.WithDescription(
				"Move a seed to a specific position in the priority list.",
			),
			mcp.WithString(
				constant.NameParameter,
				mcp.Required(),
				mcp.Description("Seed name"),
			),
			mcp.WithNumber(
				"position",
				mcp.Required(),
				mcp.Description("Target position (1 = top)"),
			),
		),
		mcp.NewTypedToolHandler(s.reorderSeed),
	)
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
