package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/timeline"
)

func (s *Server) GetTimeline(
	_ context.Context,
	r server.GetTimelineRequestObject,
) (server.GetTimelineResponseObject, error) {
	a := argument.Timeline{Limit: 50}

	if r.Params.Limit != nil {
		a.Limit = *r.Params.Limit
	}

	if r.Params.Offset != nil {
		a.Offset = *r.Params.Offset
	}

	if r.Params.Since != nil {
		a.Since = *r.Params.Since
	}

	merged, timelineError := s.service.Timeline(a)

	if timelineError != nil {
		return server.GetTimeline500JSONResponse(
			*s.captureFail(timelineError, constant.UnexpectedError),
		), nil
	}

	var result []server.TimelineEntry

	for _, e := range merged {
		description := timeline.FormatDescription(e)
		entry := server.TimelineEntry{
			Timestamp: e.Timestamp,
			Kind:      e.Kind,
			Actor:     e.Actor,
			Subject:   description,
		}
		result = append(result, entry)
	}

	if result == nil {
		result = []server.TimelineEntry{}
	}

	return server.GetTimeline200JSONResponse(result), nil
}
