package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerContainerSnapshots() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListContainerSnapshots,
			mcp.WithDescription("List snapshots for an LXC container"),
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
		mcp.NewTypedToolHandler(s.ListContainerSnapshots),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateContainerSnapshot,
			mcp.WithDescription("Create a snapshot of an LXC container"),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Container ID"),
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
		mcp.NewTypedToolHandler(s.CreateContainerSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RollbackContainerSnapshot,
			mcp.WithDescription("Rollback an LXC container to a snapshot"),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Container ID"),
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
		mcp.NewTypedToolHandler(s.RollbackContainerSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteContainerSnapshot,
			mcp.WithDescription("Delete a snapshot from an LXC container"),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Container ID"),
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
		mcp.NewTypedToolHandler(s.DeleteContainerSnapshot),
	)
}
