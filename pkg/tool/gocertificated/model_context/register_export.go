package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerExport() {
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
