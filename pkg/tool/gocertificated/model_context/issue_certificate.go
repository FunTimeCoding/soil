package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) issueCertificate(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.IssueCertificate,
) (*mcp.CallToolResult, error) {
	b := &server.CertificateBody{
		Authority:  a.Authority,
		Kind:       server.LeafKind(a.Kind),
		CommonName: a.CommonName,
	}
	optionalSlice(&b.Host, a.Host)
	optionalNumber(&b.ValidDay, a.ValidDay)
	result, key, e := s.service.IssueCertificate(b)

	if e != nil {
		return s.captureFail(e, constant.IssueFail)
	}

	certificate := convert.Certificate(result)
	certificate.PrivateKey = &key

	return response.SuccessAny(certificate)
}
