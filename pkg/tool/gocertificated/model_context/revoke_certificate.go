package model_context

import (
	"context"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) revokeCertificate(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Serial,
) (*mcp.CallToolResult, error) {
	existing, e := s.store.BySerial(a.Serial)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	if existing == nil || existing.Revoked != nil {
		return response.Fail("%s", constant.CertificateMissing)
	}

	if f := s.store.Revoke(a.Serial, time.Now()); f != nil {
		return s.captureFail(f, constant.RevokeFail)
	}

	result, g := s.store.BySerial(a.Serial)

	if g != nil {
		return s.captureFail(g, library.UnexpectedError)
	}

	return response.SuccessAny(convert.Certificate(result))
}
