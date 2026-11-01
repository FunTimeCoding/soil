package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerWriting() {
	s.server.AddTool(
		mcp.NewTool(
			constant.SaveMemory,
			mcp.WithDescription(
				"Create a new memory. Returns the created memory ID and its stored tags.",
			),
			mcp.WithString(
				constant.MemoryName,
				mcp.Required(),
				mcp.Description(
					"Short name for identification (e.g. 'retry policy', 'careful work')",
				),
			),
			mcp.WithString(
				constant.Content,
				mcp.Required(),
				mcp.Description("The memory content"),
			),
			mcp.WithString(
				constant.Description,
				mcp.Required(),
				mcp.Description("One-line description for index and search"),
			),
			mcp.WithString(
				constant.Type,
				mcp.Description(
					"Memory type: user, feedback, project, reference (default: feedback)",
				),
			),
			mcp.WithString(
				constant.Source,
				mcp.Description(
					"Session name of the caller (optional, for version attribution)",
				),
			),
			mcp.WithNumber(
				constant.ParentIdentifier,
				mcp.Description("Parent memory ID (optional, for hierarchy)"),
			),
			mcp.WithString(
				constant.Tags,
				mcp.Description(
					"Comma-separated tags to apply at creation, in the same transaction as the memory. Special tags: 'always' loads the memory in full on every profile call, 'no-index' hides it from the profile index.",
				),
			),
			mcp.WithString(
				constant.Base,
				mcp.Description(constant.BaseParameterDescription),
			),
		),
		s.create,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateMemory,
			mcp.WithDescription(constant.UpdateMemoryDescription),
			mcp.WithNumber(
				constant.MemoryIdentifier,
				mcp.Required(),
				mcp.Description("ID of the memory to update"),
			),
			mcp.WithString(
				constant.MemoryName,
				mcp.Description("New name; omit to keep the current one"),
			),
			mcp.WithString(
				constant.Content,
				mcp.Description("New content; omit to keep the current body"),
			),
			mcp.WithString(
				constant.Description,
				mcp.Description(
					"New one-line description for index and search; omit to keep the current one",
				),
			),
			mcp.WithString(
				constant.Source,
				mcp.Description(
					"Session name of the caller (optional, for version attribution)",
				),
			),
			mcp.WithString(
				constant.Base,
				mcp.Description(constant.BaseUpdateDescription),
			),
			mcp.WithBoolean(
				constant.ClearBase,
				mcp.Description(constant.ClearBaseDescription),
			),
		),
		s.update,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ForgetMemory,
			mcp.WithDescription(
				"Soft-delete a memory. The memory is marked inactive and a version record is created. Content is preserved in history.",
			),
			mcp.WithNumber(
				constant.MemoryIdentifier,
				mcp.Required(),
				mcp.Description("Memory ID to forget"),
			),
		),
		s.forget,
	)
}
