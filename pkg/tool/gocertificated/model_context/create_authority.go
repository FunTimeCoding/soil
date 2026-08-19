package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createAuthority(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateAuthority,
) (*mcp.CallToolResult, error) {
	b := &server.AuthorityBody{
		Name:       a.Name,
		Kind:       server.AuthorityKind(a.Kind),
		CommonName: a.CommonName,
	}
	optionalText(&b.Country, a.Country)
	optionalText(&b.Province, a.Province)
	optionalText(&b.Organization, a.Organization)
	optionalSlice(&b.PermittedDomain, a.PermittedDomain)
	optionalSlice(&b.PermittedAddress, a.PermittedAddress)
	optionalNumber(&b.ValidYear, a.ValidYear)
	result, e := s.service.CreateAuthority(b)

	if errors.Is(e, constant.ErrorConflict) {
		return response.Fail("%s", constant.AuthorityLive)
	}

	if e != nil {
		return s.captureFail(e, constant.CreateAuthorityFail)
	}

	return response.SuccessAny(convert.Authority(result))
}
