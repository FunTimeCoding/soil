package tool

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"slices"
)

func CreateMachine() mcp.Tool {
	return mcp.NewTool(
		constant.CreateMachine,
		slices.Concat(
			[]mcp.ToolOption{
				mcp.WithDescription(
					"Create a virtual machine. Waits for the creation task to complete before returning.",
				),
				mcp.WithString(
					constant.NodeParameter,
					mcp.Required(),
					mcp.Description("Target node name"),
				),
				mcp.WithString(
					generativeConstant.ParameterName,
					mcp.Required(),
					mcp.Description("VM name"),
				),
				mcp.WithNumber(
					generativeConstant.ParameterIdentifier,
					mcp.Description("VMID (auto-allocated when omitted)"),
				),
			},
			hardwareOptions(),
			diskOptions(),
			cloudInitOptions(),
			managementOptions(),
		)...,
	)
}
