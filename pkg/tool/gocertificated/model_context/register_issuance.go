package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerIssuance() {
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
}
