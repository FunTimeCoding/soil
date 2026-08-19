package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) listCertificates(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListCertificates,
) (*mcp.CallToolResult, error) {
	f := store.NewFilter()
	f.Authority = a.Authority
	f.Kind = a.Kind
	f.Limit = int(a.Limit)

	if a.Revoked {
		f.Revoked = &a.Revoked
	}

	if a.ExpiresBefore != "" {
		before, e := time.Parse(time.RFC3339, a.ExpiresBefore)

		if e != nil {
			return response.Fail("expires_before is not a timestamp: %v", e)
		}

		f.Before = &before
	}

	result, e := s.store.Certificates(f)

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.SuccessAny(convert.Certificates(result))
}
