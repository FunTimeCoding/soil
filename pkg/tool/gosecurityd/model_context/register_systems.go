package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerSystems() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSystems,
			mcp.WithDescription(
				"List systems with their tracker coverage and finding counts, by priority. A system is clean only on a supported release with a loaded tracker, a fresh snapshot, a recently verified tracker and no finding; on any other coverage an empty finding list means unchecked, not clean. stale marks a snapshot past its age threshold, tracker_stale a tracker that has not answered within its threshold.",
			),
		),
		mcp.NewTypedToolHandler(s.ListSystems),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetSystemFindings,
			mcp.WithDescription(
				"Get one system's tracker coverage and findings: per CVE the source package, installed and fixed version, where the fix is (debian, paid support only, no fix) and whether the upgrade candidate already has it",
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("System name"),
			),
		),
		mcp.NewTypedToolHandler(s.GetSystemFindings),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListFindings,
			mcp.WithDescription(
				"List findings grouped by CVE with the systems each one hits, most systems first",
			),
		),
		mcp.NewTypedToolHandler(s.ListFindings),
	)
}
