package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/tool"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerMachineConfiguration() {
	s.server.AddTool(
		tool.CreateMachine(),
		mcp.NewTypedToolHandler(s.CreateMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateMachine,
			mcp.WithDescription(
				"Update configuration of an existing virtual machine.",
			),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to search all nodes)"),
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Description("New VM name"),
			),
			mcp.WithString(
				constant.TagsOption,
				mcp.Description("Semicolon-separated tags (replaces existing)"),
			),
			mcp.WithBoolean(
				constant.OnBootOption,
				mcp.Description("Start VM automatically on host boot"),
			),
			mcp.WithNumber(constant.CoresOption, mcp.Description("CPU cores")),
			mcp.WithNumber(
				constant.MemoryOption,
				mcp.Description("Memory in MiB"),
			),
			mcp.WithString(
				constant.DescriptionOption,
				mcp.Description("VM description"),
			),
			mcp.WithString(
				constant.DeleteOption,
				mcp.Description(
					"Comma-separated field names to clear, e.g. tags,description",
				),
			),
		),
		mcp.NewTypedToolHandler(s.UpdateMachine),
	)
}
