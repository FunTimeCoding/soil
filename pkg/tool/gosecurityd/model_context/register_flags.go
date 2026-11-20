package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerFlags() {
	s.server.AddTool(
		mcp.NewTool(
			constant.FlagVulnerability,
			mcp.WithDescription(
				"Flag a CVE as a concern to mitigate. Stores the NVD severity and score as they are now; a CVE NVD has not published yet is flagged without them. Flagging a mitigated CVE again reopens it.",
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("CVE identifier, e.g. CVE-2024-3094"),
			),
			mcp.WithString(
				constant.NoteField,
				mcp.Description("Why it is a concern, e.g. which hosts run it"),
			),
		),
		mcp.NewTypedToolHandler(s.FlagVulnerability),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListFlags,
			mcp.WithDescription(
				"List flagged CVEs, newest first. Open concerns only unless include_mitigated is set.",
			),
			mcp.WithBoolean(
				"include_mitigated",
				mcp.Description("Also list flags already mitigated"),
			),
		),
		mcp.NewTypedToolHandler(s.ListFlags),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.MitigateFlag,
			mcp.WithDescription(
				"Mark a flagged CVE as mitigated, with how it was mitigated",
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("CVE identifier, e.g. CVE-2024-3094"),
			),
			mcp.WithString(
				"mitigation",
				mcp.Description(
					"How it was mitigated, e.g. upgraded to 5.6.1-1",
				),
			),
		),
		mcp.NewTypedToolHandler(s.MitigateFlag),
	)
}
