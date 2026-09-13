package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetSessions(
	_ context.Context,
	r server.GetSessionsRequestObject,
) (server.GetSessionsResponseObject, error) {
	limit := 0

	if r.Params.Limit != nil {
		limit = *r.Params.Limit
	}

	offset := 0

	if r.Params.Offset != nil {
		offset = *r.Params.Offset
	}

	peek := r.Params.Peek != nil && *r.Params.Peek
	list := s.service.EnrichedSessions

	if r.Params.Unnamed != nil && *r.Params.Unnamed {
		list = s.service.UnnamedSessions
	}

	sessions, e := list(limit, offset)

	if e != nil {
		return server.GetSessions500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	identifiers := make([]string, 0, len(sessions))

	for _, i := range sessions {
		identifiers = append(identifiers, i.Identifier)
	}

	labels, f := s.service.LabelsBySessions(identifiers)

	if f != nil {
		return server.GetSessions500JSONResponse(
			*s.captureFail(f, constant.UnexpectedError),
		), nil
	}

	var result []server.SessionDetail

	for _, i := range sessions {
		d := server.SessionDetail{
			Identifier: i.Identifier,
			Timestamp:  i.Timestamp,
			Lines:      i.Lines,
		}

		if i.Name != "" {
			d.Name = &i.Name
		}

		if i.Slug != "" {
			d.Slug = &i.Slug
		}

		if i.WorkDirectory != "" {
			d.WorkDirectory = &i.WorkDirectory
		}

		if i.Branch != "" {
			d.Branch = &i.Branch
		}

		if i.Alias != "" {
			d.Alias = &i.Alias
		}

		if i.Description != "" {
			d.Description = &i.Description
		}

		if entries := labels[i.Identifier]; len(entries) > 0 {
			carried := make([]server.LabelEntry, 0, len(entries))

			for _, entry := range entries {
				carried = append(
					carried,
					server.LabelEntry{Key: entry.Key, Value: entry.Value},
				)
			}

			d.Labels = &carried
		}

		if peek {
			preview := s.service.FirstUserMessage(i.Identifier)

			if len(preview) > 60 {
				preview = preview[:60]
			}

			if preview != "" {
				d.Preview = &preview
			}
		}

		result = append(result, d)
	}

	return server.GetSessions200JSONResponse{Sessions: result}, nil
}
