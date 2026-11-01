package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerConversations() {
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchConversations,
			mcp.WithDescription(
				"Search every Claude Code conversation in the harbor for literal text - case-insensitive substrings, not meaning (goqueryd searches by meaning). A conversation matches when each term appears somewhere in it; results group by conversation, newest hit first, with up to three snippets each, blocks holding every term first. Searches messages by default; kinds adds edits and tool calls. Tool results are not searched. Read around a hit with read_conversation, passing the conversation's session and the hit's identifier.",
			),
			mcp.WithString(
				generativeConstant.ParameterQuery,
				mcp.Required(),
				mcp.Description(
					"Terms separated by spaces, each matched as a substring",
				),
			),
			mcp.WithArray(
				constant.Kinds,
				mcp.WithStringItems(),
				mcp.Description(
					"Block kinds to search: message, edit, call (default: message)",
				),
			),
			mcp.WithNumber(
				constant.Limit,
				mcp.Description(
					"Maximum conversations (default 20, at most 100)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.searchConversations),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ReadConversation,
			mcp.WithDescription(
				"Read the blocks around one search hit in a conversation - messages, edits and tool calls in order, each with time, role, kind and identifier, long blocks cut at 2000 characters. Tool results are not included.",
			),
			mcp.WithString(
				constant.Session,
				mcp.Required(),
				mcp.Description(
					"Conversation session identifier from search_conversations",
				),
			),
			mcp.WithString(
				constant.Around,
				mcp.Required(),
				mcp.Description("Block identifier of the hit to read around"),
			),
			mcp.WithNumber(
				constant.Count,
				mcp.Description(
					"Blocks on each side of the hit (default 5, at most 25)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.readConversation),
	)
}
