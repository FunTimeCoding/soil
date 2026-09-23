package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func (s *Server) GetTokens(
	_ context.Context,
	r server.GetTokensRequestObject,
) (server.GetTokensResponseObject, error) {
	scope := ""

	if r.Params.Scope != nil {
		scope = *r.Params.Scope
	}

	summary, e := s.service.MemoryTokenSummary(scope)

	if e != nil {
		return server.GetTokens500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	response := server.GetTokens200JSONResponse{
		Statistic: make([]server.TokenStatistic, 0, len(summary.Statistic)),
		Block:     spread(summary.Block),
		Description: spread(summary.Description),
		Withheld:    summary.Withheld,
	}

	for _, t := range summary.Statistic {
		if t.Hidden {
			continue
		}

		entry := server.TokenStatistic{
			Identifier:  int64(t.Identifier),
			Name:        t.Name,
			Block:       t.Block,
			Description: t.Description,
		}

		if len(t.Tags) > 0 {
			entry.Tags = &t.Tags
		}

		response.Statistic = append(response.Statistic, entry)
	}

	return response, nil
}
