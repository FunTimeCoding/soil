package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
	"slices"
)

func (s *Server) GetRedactedMemories(
	_ context.Context,
	_ server.GetRedactedMemoriesRequestObject,
) (server.GetRedactedMemoriesResponseObject, error) {
	identifiers, e := s.service.HiddenIdentifiers()

	if e != nil {
		return server.GetRedactedMemories500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	result := make([]int64, 0, len(identifiers))

	for identifier := range identifiers {
		result = append(result, identifier)
	}

	slices.Sort(result)

	return server.GetRedactedMemories200JSONResponse(result), nil
}
