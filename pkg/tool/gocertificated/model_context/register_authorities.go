package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerAuthorities() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListAuthorities,
			mcp.WithDescription(
				"List every certificate authority in the chain, root first.",
			),
		),
		mcp.NewTypedToolHandler(s.listAuthorities),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateAuthority,
			mcp.WithDescription(
				"Create the root or a name-constrained intermediate. Fails when an authority of that name is already live.",
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description(
					"Short slug identifying the authority - root, cluster, host.",
				),
			),
			mcp.WithString(
				constant.KindParameter,
				mcp.Required(),
				mcp.Description("root or intermediate."),
			),
			mcp.WithString(
				constant.CommonNameParameter,
				mcp.Required(),
				mcp.Description(
					"Subject common name as it appears in certificate viewers.",
				),
			),
			mcp.WithString(
				constant.CountryParameter,
				mcp.Description(
					"Two-letter country. Required for the root, inherited by intermediates.",
				),
			),
			mcp.WithString(
				constant.ProvinceParameter,
				mcp.Description(
					"State or province. Required for the root, inherited by intermediates.",
				),
			),
			mcp.WithString(
				constant.OrganizationParameter,
				mcp.Description(
					"Organization. Required for the root, inherited by intermediates.",
				),
			),
			mcp.WithArray(
				constant.DomainParameter,
				mcp.Description(
					"Permitted DNS subtrees for an intermediate. Each entry covers itself and every name to its left.",
				),
			),
			mcp.WithArray(
				constant.AddressParameter,
				mcp.Description(
					"Permitted address ranges for an intermediate, in CIDR notation.",
				),
			),
			mcp.WithNumber(
				constant.ValidYearParameter,
				mcp.Description(
					"Validity in years (default 20 for a root, 10 for an intermediate).",
				),
			),
		),
		mcp.NewTypedToolHandler(s.createAuthority),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetAuthority,
			mcp.WithDescription("Inspect one authority by its slug."),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Authority slug."),
			),
		),
		mcp.NewTypedToolHandler(s.getAuthority),
	)
}
