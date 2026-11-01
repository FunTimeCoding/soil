package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerMachineCopies() {
	s.server.AddTool(
		mcp.NewTool(
			constant.CloneMachine,
			mcp.WithDescription(
				"Clone a virtual machine. Waits for the clone task to complete before returning.",
			),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Source machine ID to clone from"),
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Name for the new VM"),
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to search all nodes)"),
			),
			mcp.WithNumber(
				"new_identifier",
				mcp.Description(
					"VMID for the clone (auto-allocated when omitted)",
				),
			),
			mcp.WithBoolean(
				"full",
				mcp.Description(
					"Full clone instead of linked clone (default false)",
				),
			),
			mcp.WithString(
				constant.StorageParameter,
				mcp.Description("Target storage for the clone"),
			),
			mcp.WithString(
				"snapshot",
				mcp.Description(
					"Clone from a specific snapshot instead of current state",
				),
			),
		),
		mcp.NewTypedToolHandler(s.CloneMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteMachine,
			mcp.WithDescription(
				"Delete a virtual machine. The VM must be stopped first. Waits for deletion to complete.",
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
			mcp.WithBoolean(
				"purge",
				mcp.Description(
					"Remove from cluster config and associated resources (default false)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.DeleteMachine),
	)
}
