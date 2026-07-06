package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func (s *Server) GetProfile(
	_ context.Context,
	r server.GetProfileRequestObject,
) (server.GetProfileResponseObject, error) {
	topic := ""

	if r.Params.Topic != nil {
		topic = *r.Params.Topic
	}

	showDetail := r.Params.Detail != nil && *r.Params.Detail
	result, d, e := s.service.Profile(topic, showDetail)

	if e != nil {
		return server.GetProfile500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	response := server.GetProfile200JSONResponse{
		Always: make([]server.ProfileMemory, 0, len(result.Always)),
		Index:  make([]server.ProfileSummary, 0, len(result.Index)),
	}

	for _, m := range result.Always {
		entry := server.ProfileMemory{
			Identifier:  int(m.Identifier),
			Name:        m.Name,
			Content:     m.Content,
			Description: m.Description,
			Type:        m.Type,
			CreatedAt:   m.CreatedAt,
			UpdatedAt:   m.UpdatedAt,
			IsActive:    m.IsActive,
		}

		if len(m.Tags) > 0 {
			entry.Tags = &m.Tags
		}

		response.Always = append(response.Always, entry)
	}

	if len(result.Relevant) > 0 {
		relevant := make(
			[]server.ProfileSearchResult,
			0,
			len(result.Relevant),
		)

		for _, r := range result.Relevant {
			entry := server.ProfileSearchResult{
				Identifier:  int(r.Identifier),
				Name:        r.Name,
				Content:     r.Content,
				Description: r.Description,
				Type:        r.Type,
				UpdatedAt:   r.UpdatedAt,
				Rank:        float32(r.Rank),
			}

			if len(r.Tags) > 0 {
				entry.Tags = &r.Tags
			}

			relevant = append(relevant, entry)
		}

		response.Relevant = &relevant
	}

	for _, m := range result.Index {
		entry := server.ProfileSummary{
			Identifier:  int(m.Identifier),
			Name:        m.Name,
			Description: m.Description,
			Type:        m.Type,
			UpdatedAt:   m.UpdatedAt,
		}

		if len(m.Tags) > 0 {
			entry.Tags = &m.Tags
		}

		if len(m.Children) > 0 {
			entry.Children = &m.Children
		}

		response.Index = append(response.Index, entry)
	}

	if len(result.Impressions) > 0 {
		impressions := make(
			[]server.ProfileImpression,
			0,
			len(result.Impressions),
		)

		for _, i := range result.Impressions {
			impressions = append(
				impressions,
				server.ProfileImpression{
					Identifier: int(i.Identifier),
					Content:    i.Content,
					Source:     i.Source,
					CreatedAt:  i.CreatedAt,
				},
			)
		}

		response.Impressions = &impressions
	}

	if len(result.Completions) > 0 {
		completions := make(
			[]server.ProfileCompletion,
			0,
			len(result.Completions),
		)

		for _, c := range result.Completions {
			completions = append(
				completions,
				server.ProfileCompletion{
					SessionName: c.SessionName,
					Body:        c.Body,
				},
			)
		}

		response.Completions = &completions
	}

	if d != nil {
		response.Detail = &server.ProfileDetail{
			Budget:             d.Budget,
			AlwaysTokens:       d.AlwaysTokens,
			IndexTokens:        d.IndexTokens,
			IndexTrimmed:       d.IndexTrimmed,
			CompletionTokens:   d.CompletionTokens,
			CompletionsTrimmed: d.CompletionsTrimmed,
			ImpressionTokens:   d.ImpressionTokens,
			ImpressionsTrimmed: d.ImpressionsTrimmed,
			RelevantTokens:     d.RelevantTokens,
			RelevantTrimmed:    d.RelevantTrimmed,
			TotalTokens:        d.TotalTokens,
		}
	}

	return response, nil
}
