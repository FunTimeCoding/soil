package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerContainers() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListContainers,
			mcp.WithDescription(
				"List LXC containers, optionally filtered by node",
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to list across all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ListContainers),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetContainer,
			mcp.WithDescription(
				"Get status and config for an LXC container by VMID",
			),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Container ID"),
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.GetContainer),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.StartContainer,
			mcp.WithDescription("Start an LXC container"),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Container ID"),
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.StartContainer),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ShutdownContainer,
			mcp.WithDescription(
				"Gracefully shutdown an LXC container - prefer this over stop",
			),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Container ID"),
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ShutdownContainer),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.StopContainer,
			mcp.WithDescription(
				"Stop an LXC container immediately - a hard kill, use shutdown_container unless the guest is unresponsive",
			),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Container ID"),
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.StopContainer),
	)
}
