package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerSnippets() {
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateSnippet,
			mcp.WithDescription(
				"Create a cloud-init or other snippet file on the Proxmox host via SSH. Returns the volume reference for use with cicustom.",
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Snippet filename, e.g. cloud-init-agent.yml"),
			),
			mcp.WithString(
				constant.ContentParameter,
				mcp.Required(),
				mcp.Description("File content, e.g. a cloud-init YAML body"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateSnippet),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetSnippet,
			mcp.WithDescription("Read a snippet file from the Proxmox host"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Snippet filename"),
			),
		),
		mcp.NewTypedToolHandler(s.GetSnippet),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSnippets,
			mcp.WithDescription("List snippet files on the Proxmox host"),
		),
		mcp.NewTypedToolHandler(s.ListSnippets),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteSnippet,
			mcp.WithDescription("Delete a snippet file from the Proxmox host"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Snippet filename"),
			),
		),
		mcp.NewTypedToolHandler(s.DeleteSnippet),
	)
}
