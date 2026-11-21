package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerAdvisory() {
	s.server.AddTool(
		mcp.NewTool(
			constant.GetVulnerability,
			mcp.WithDescription(
				"Get a vulnerability from the National Vulnerability Database: description, score, severity, vector and references.",
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("CVE identifier, e.g. CVE-2024-3094"),
			),
		),
		mcp.NewTypedToolHandler(s.GetVulnerability),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetDebianStatus,
			mcp.WithDescription(
				"Get Debian's fix status for a vulnerability per source package and release: status, fixed version (never_affected when the release never shipped the flaw), urgency and the version per archive. Use it before judging a Debian host by version ranges.",
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("CVE identifier, e.g. CVE-2024-3094"),
			),
			mcp.WithString(
				constant.ReleaseParameter,
				mcp.Description(
					"Debian release codename to filter by, e.g. trixie",
				),
			),
		),
		mcp.NewTypedToolHandler(s.GetDebianStatus),
	)
}
