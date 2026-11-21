package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerRemediation() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListRemediation,
			mcp.WithDescription(
				"List the remediation for vulnerabilities at or above a score: upgrades, release upgrades, reboots, kernels out of reach and vulnerabilities without an upgrade, each with the subject decisions are keyed by; open items first. How items form: gosecurityd://guide/remediation.",
			),
			mcp.WithNumber(
				constant.ScoreField,
				mcp.Description(
					"The lowest score a finding needs to count; empty means 9.0 (critical)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.ListRemediation),
	)
}
