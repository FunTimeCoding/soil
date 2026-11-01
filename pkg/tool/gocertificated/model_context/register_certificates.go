package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerCertificates() {
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
				generativeConstant.ParameterLimit,
				mcp.Description("Maximum number of results."),
			),
		),
		mcp.NewTypedToolHandler(s.listCertificates),
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
}
