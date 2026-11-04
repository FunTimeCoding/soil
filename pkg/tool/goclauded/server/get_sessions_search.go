package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetSessionsSearch(
	x context.Context,
	r server.GetSessionsSearchRequestObject,
) (server.GetSessionsSearchResponseObject, error) {
	var kinds []string

	if r.Params.Kinds != nil {
		kinds = *r.Params.Kinds
	}

	limit := 0

	if r.Params.Limit != nil {
		limit = *r.Params.Limit
	}

	found, e := s.search.SearchConversations(r.Params.Query, kinds, limit)

	if e != nil {
		return server.GetSessionsSearch500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	conversations := []server.SearchConversation{}

	for _, c := range found {
		conversations = append(conversations, searchConversation(c))
	}

	indexed, total := s.search.SearchProgress()

	return server.GetSessionsSearch200JSONResponse{
		Conversations: conversations,
		Indexed:       int(indexed),
		Total:         int(total),
	}, nil
}
