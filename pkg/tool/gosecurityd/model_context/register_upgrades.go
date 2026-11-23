package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerUpgrades() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListUpgrades,
			append(
				selectionOptions(),
				mcp.WithDescription(
					"List the package upgrades that fix findings on the selected systems, by release and source package, with their hosts, versions, vulnerabilities and top score; the most hosts times vulnerabilities first, the kernel apart.",
				),
				mcp.WithString(
					constant.ReleaseParameter,
					mcp.Description(
						"A release codename to narrow to, e.g. bookworm; empty lists every release",
					),
				),
				mcp.WithBoolean(
					constant.KernelParameter,
					mcp.Description(
						"True lists the kernel upgrades instead of the package upgrades",
					),
				),
				mcp.WithNumber(
					constant.ScoreField,
					mcp.Description(constant.ScoreDescription),
				),
			)...,
		),
		mcp.NewTypedToolHandler(s.ListUpgrades),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetUpgrade,
			mcp.WithDescription(
				"Get one package upgrade with its hosts and, per host, the binary packages to pass to apt-get install --only-upgrade.",
			),
			mcp.WithString(
				constant.ReleaseParameter,
				mcp.Required(),
				mcp.Description("Debian release codename, e.g. bookworm"),
			),
			mcp.WithString(
				constant.SourceParameter,
				mcp.Required(),
				mcp.Description("Source package name, e.g. openssl"),
			),
		),
		mcp.NewTypedToolHandler(s.GetUpgrade),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetScoreStatus,
			mcp.WithDescription(
				"Get how far the score mirror has come: whether the first fill is complete, how many vulnerabilities carry a score, and when it last caught up.",
			),
		),
		mcp.NewTypedToolHandler(s.GetScoreStatus),
	)
}
