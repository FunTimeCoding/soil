package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func (s *Server) GetVersions(
	_ context.Context,
	r server.GetVersionsRequestObject,
) (server.GetVersionsResponseObject, error) {
	limit := 50

	if r.Params.Limit != nil {
		limit = *r.Params.Limit
	}

	offset := 0

	if r.Params.Offset != nil {
		offset = *r.Params.Offset
	}

	versions, e := s.service.VersionsSince(r.Params.Since, limit, offset)

	if e != nil {
		return server.GetVersions500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	entries := make([]server.VersionEntry, 0, len(versions))

	for _, v := range versions {
		entries = append(
			entries,
			server.VersionEntry{
				Identifier:       int(v.Identifier),
				MemoryIdentifier: int(v.MemoryIdentifier),
				Name:             v.Name,
				Description:      v.Description,
				ChangedAt:        v.ChangedAt,
				ChangeType:       v.ChangeType,
				Source:           v.Source,
			},
		)
	}

	return server.GetVersions200JSONResponse(entries), nil
}
