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
				"List systems by priority: coverage, clean, package finding counts, critical (vulnerabilities at 9.0 or higher), top score and kernel state. Coverage, clean and kernel states: gosecurityd://guide/matching.",
			),
			mcp.WithNumber(
				constant.ScoreField,
				mcp.Description(constant.ScoreDescription),
			),
		),
		mcp.NewTypedToolHandler(s.ListSystems),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetSystemFindings,
			mcp.WithDescription(
				"Get one system's coverage and findings: per vulnerability the source package, installed and fixed version, where the fix is, whether the upgrade candidate has it, and the score.",
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
				"List vulnerabilities across systems with their score and how many systems each fix kind covers, the highest score first; unimportant marks one Debian will not fix.",
			),
			mcp.WithNumber(
				constant.ScoreField,
				mcp.Description(constant.ScoreDescription),
			),
		),
		mcp.NewTypedToolHandler(s.ListFindings),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetAffectedSystems,
			mcp.WithDescription(
				"Get the systems one vulnerability hits, each with its source package, versions, upgrade candidate and fix kind.",
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("CVE identifier, e.g. CVE-2024-3094"),
			),
		),
		mcp.NewTypedToolHandler(s.GetAffectedSystems),
	)
}
