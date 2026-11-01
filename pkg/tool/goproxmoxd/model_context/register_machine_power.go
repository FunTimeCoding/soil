package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerMachinePower() {
	s.server.AddTool(
		mcp.NewTool(
			constant.StartMachine,
			mcp.WithDescription("Start a virtual machine"),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.StartMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.StopMachine,
			mcp.WithDescription("Stop a virtual machine immediately"),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.StopMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ShutdownMachine,
			mcp.WithDescription("Gracefully shutdown a virtual machine"),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ShutdownMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ResetMachine,
			mcp.WithDescription("Reset a virtual machine"),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ResetMachine),
	)
}
