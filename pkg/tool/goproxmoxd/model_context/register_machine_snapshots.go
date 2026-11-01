package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerMachineSnapshots() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListMachineSnapshots,
			mcp.WithDescription("List snapshots for a virtual machine"),
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
		mcp.NewTypedToolHandler(s.ListMachineSnapshots),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateMachineSnapshot,
			mcp.WithDescription("Create a snapshot of a virtual machine"),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Snapshot name"),
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateMachineSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RollbackMachineSnapshot,
			mcp.WithDescription("Rollback a virtual machine to a snapshot"),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Snapshot name"),
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.RollbackMachineSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteMachineSnapshot,
			mcp.WithDescription("Delete a snapshot from a virtual machine"),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Snapshot name"),
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.DeleteMachineSnapshot),
	)
}
