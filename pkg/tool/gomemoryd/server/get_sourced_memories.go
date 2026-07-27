package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func (s *Server) GetSourcedMemories(
	_ context.Context,
	r server.GetSourcedMemoriesRequestObject,
) (server.GetSourcedMemoriesResponseObject, error) {
	memories, e := s.service.ListDocumentSourced(r.Params.Scope)

	if e != nil {
		return server.GetSourcedMemories500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	result := make([]server.SourcedMemory, 0, len(memories))

	for _, m := range memories {
		result = append(
			result,
			server.SourcedMemory{
				Identifier:       m.Identifier,
				Name:             m.Name,
				ParentIdentifier: m.ParentIdentifier,
				ProvenanceFile:   m.ProvenanceFile,
				ProvenanceAnchor: m.ProvenanceAnchor,
				ProvenanceHash:   m.ProvenanceHash,
				Ordinal:          m.Ordinal,
			},
		)
	}

	return server.GetSourcedMemories200JSONResponse(result), nil
}
