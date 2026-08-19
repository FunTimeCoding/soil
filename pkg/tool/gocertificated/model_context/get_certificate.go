package model_context

import (
	"context"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getCertificate(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Serial,
) (*mcp.CallToolResult, error) {
	result, e := s.store.BySerial(a.Serial)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	if result == nil {
		return response.Fail("%s", constant.CertificateMissing)
	}

	return response.SuccessAny(convert.Certificate(result))
}
