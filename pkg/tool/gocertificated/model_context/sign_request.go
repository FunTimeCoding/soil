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

func (s *Server) signRequest(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SignRequest,
) (*mcp.CallToolResult, error) {
	b := &server.SigningRequestBody{
		Authority: a.Authority,
		Kind:      server.LeafKind(a.Kind),
		Request:   a.Request,
	}
	optionalNumber(&b.ValidDay, a.ValidDay)
	result, e := s.service.SignRequest(b)

	if e != nil {
		return s.captureFail(e, constant.SignFail)
	}

	return response.SuccessAny(convert.Certificate(result))
}
