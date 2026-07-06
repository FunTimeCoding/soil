package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetSessionDetail(
	_ context.Context,
	r server.GetSessionDetailRequestObject,
) (server.GetSessionDetailResponseObject, error) {
	d, e := s.service.SessionDetail(r.Identifier)

	if e != nil {
		return server.GetSessionDetail500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	if d == nil {
		return server.GetSessionDetail404JSONResponse{
			Error: fmt.Sprintf("session not found: %s", r.Identifier),
		}, nil
	}

	result := server.SessionDetailResponse{Identifier: d.Identifier}

	if d.Session != nil {
		if d.Session.Name != "" {
			result.Name = &d.Session.Name
		}

		if d.Session.Callsign != nil {
			result.Callsign = d.Session.Callsign
		}
	}

	if d.Slug != "" {
		result.Slug = &d.Slug
	}

	if d.Created != "" {
		result.Created = &d.Created
	}

	if d.TurnCount > 0 {
		result.TurnCount = &d.TurnCount
	}

	if d.Alias != "" {
		result.Alias = &d.Alias
	}

	if d.Description != "" {
		result.Description = &d.Description
	}

	if len(d.Completions) > 0 {
		var entries []server.SessionDetailCompletion

		for _, c := range d.Completions {
			entry := server.SessionDetailCompletion{
				Kind:  c.Kind,
				Topic: c.Topic,
			}

			if c.Summary != "" {
				entry.Summary = &c.Summary
			}

			ts := c.CreatedAt.Format("2006-01-02T15:04:05Z")
			entry.Timestamp = &ts
			entries = append(entries, entry)
		}

		result.Completions = &entries
	}

	if d.Summary != "" {
		result.Summary = &d.Summary
	}

	labels, labelError := s.service.LabelsBySession(d.Identifier)

	if labelError != nil {
		return server.GetSessionDetail500JSONResponse(
			*s.captureFail(labelError, constant.UnexpectedError),
		), nil
	}

	if len(labels) > 0 {
		var le []server.LabelEntry

		for _, l := range labels {
			le = append(
				le,
				server.LabelEntry{Key: l.Key, Value: l.Value},
			)
		}

		result.Labels = &le
	}

	pulses, pulseError := s.service.PulsesBySession(d.Identifier)

	if pulseError != nil {
		return server.GetSessionDetail500JSONResponse(
			*s.captureFail(pulseError, constant.UnexpectedError),
		), nil
	}

	if len(pulses) > 0 {
		var pe []server.PulseEntry

		for _, p := range pulses {
			entry := server.PulseEntry{
				Body:      p.Body,
				Timestamp: p.CreatedAt.Format("2006-01-02T15:04:05Z"),
			}

			if p.FromName != "" {
				entry.From = &p.FromName
			}

			pe = append(pe, entry)
		}

		result.Pulses = &pe
	}

	return server.GetSessionDetail200JSONResponse(result), nil
}
