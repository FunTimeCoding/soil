package model_context

import (
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
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
				generative.ParameterName,
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
				generative.ParameterName,
				mcp.Required(),
				mcp.Description("Authority slug."),
			),
		),
		mcp.NewTypedToolHandler(s.getAuthority),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListCertificates,
			mcp.WithDescription(
				"List issued certificates, newest expiry last. Use expires_before to find what needs renewing.",
			),
			mcp.WithString(
				constant.AuthorityParameter,
				mcp.Description("Only certificates issued by this authority."),
			),
			mcp.WithString(
				constant.KindParameter,
				mcp.Description("root, intermediate, server or client."),
			),
			mcp.WithString(
				constant.ExpiresBeforeParameter,
				mcp.Description(
					"Only certificates expiring before this RFC3339 timestamp.",
				),
			),
			mcp.WithBoolean(
				constant.RevokedParameter,
				mcp.Description("Only revoked certificates."),
			),
			mcp.WithNumber(
				generative.ParameterLimit,
				mcp.Description("Maximum number of results."),
			),
		),
		mcp.NewTypedToolHandler(s.listCertificates),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.IssueCertificate,
			mcp.WithDescription(
				"Issue a certificate and generate its key. The key is returned once and never stored - prefer sign_request when the subject can generate its own.",
			),
			mcp.WithString(
				constant.AuthorityParameter,
				mcp.Required(),
				mcp.Description("Issuing authority slug."),
			),
			mcp.WithString(
				constant.KindParameter,
				mcp.Required(),
				mcp.Description("server or client."),
			),
			mcp.WithString(
				constant.CommonNameParameter,
				mcp.Required(),
				mcp.Description("Subject common name."),
			),
			mcp.WithArray(
				constant.HostParameter,
				mcp.Description(
					"Subject alternative names - host names and addresses the certificate is valid for.",
				),
			),
			mcp.WithNumber(
				constant.ValidDayParameter,
				mcp.Description("Validity in days (default 90)."),
			),
		),
		mcp.NewTypedToolHandler(s.issueCertificate),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetCertificate,
			mcp.WithDescription("Inspect one certificate by serial."),
			mcp.WithString(
				constant.SerialParameter,
				mcp.Required(),
				mcp.Description("Hexadecimal serial."),
			),
		),
		mcp.NewTypedToolHandler(s.getCertificate),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SignRequest,
			mcp.WithDescription(
				"Sign a certificate signing request. No private key is returned or held - the subject keeps the key it generated.",
			),
			mcp.WithString(
				constant.AuthorityParameter,
				mcp.Required(),
				mcp.Description("Issuing authority slug."),
			),
			mcp.WithString(
				constant.KindParameter,
				mcp.Required(),
				mcp.Description("server or client."),
			),
			mcp.WithString(
				constant.RequestParameter,
				mcp.Required(),
				mcp.Description("Armored certificate signing request."),
			),
			mcp.WithNumber(
				constant.ValidDayParameter,
				mcp.Description("Validity in days (default 90)."),
			),
		),
		mcp.NewTypedToolHandler(s.signRequest),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RevokeCertificate,
			mcp.WithDescription(
				"Revoke a certificate. The record is kept and appears in the revocation list of its issuer.",
			),
			mcp.WithString(
				constant.SerialParameter,
				mcp.Required(),
				mcp.Description("Hexadecimal serial."),
			),
		),
		mcp.NewTypedToolHandler(s.revokeCertificate),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.PendingPublication,
			mcp.WithDescription(
				"Show which files a publish would commit to the repository.",
			),
		),
		mcp.NewTypedToolHandler(s.pendingPublication),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Publish,
			mcp.WithDescription(
				"Commit every unpublished authority to the repository in one commit.",
			),
		),
		mcp.NewTypedToolHandler(s.publish),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RootCertificate,
			mcp.WithDescription(
				"The root certificate, armored, for importing into a trust store.",
			),
		),
		mcp.NewTypedToolHandler(s.rootCertificate),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RevocationList,
			mcp.WithDescription(
				"The revocation list signed by one authority, armored.",
			),
			mcp.WithString(
				constant.AuthorityParameter,
				mcp.Required(),
				mcp.Description("Authority slug."),
			),
		),
		mcp.NewTypedToolHandler(s.revocationList),
	)
}
