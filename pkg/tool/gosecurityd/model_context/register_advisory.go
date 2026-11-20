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
				"Get a CVE from the National Vulnerability Database: description, CVSS score, severity, vector and reference links",
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
				"Get Debian's fix status for a CVE per source package and release: status, fixed version (never_affected when the release never shipped the flaw), urgency and the current version per archive. Covers only releases the Debian security tracker still lists.",
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("CVE identifier, e.g. CVE-2024-3094"),
			),
			mcp.WithString(
				"release",
				mcp.Description(
					"Debian release codename to filter by, e.g. trixie",
				),
			),
		),
		mcp.NewTypedToolHandler(s.GetDebianStatus),
	)
}
