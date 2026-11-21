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
				"Record a decision: flag a vulnerability or a remediation item as a concern, with a note. Flagging a mitigated one reopens it. Decisions: gosecurityd://guide/remediation.",
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description(
					"CVE identifier, e.g. CVE-2024-3094, or a remediation item subject from list_remediation, e.g. upgrade:bookworm:openssl",
				),
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
				"List decisions, newest first; open ones unless include_mitigated is set.",
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
				"Close a decision with how: a mitigation or an accepted risk with its reason. A mitigated remediation item counts as addressed.",
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description(
					"CVE identifier or remediation item subject, as flagged",
				),
			),
			mcp.WithString(
				"mitigation",
				mcp.Description(
					"How, e.g. upgraded to 5.6.1-1, or accepted risk: retires in the first quarter",
				),
			),
		),
		mcp.NewTypedToolHandler(s.MitigateFlag),
	)
}
