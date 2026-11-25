package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/prometheus/common/model"
	"time"
)

func (s *Server) GetRoster(
	x context.Context,
	r server.GetRosterRequestObject,
) (server.GetRosterResponseObject, error) {
	var window time.Duration

	if r.Params.SeenWithin != nil {
		d, e := model.ParseDuration(*r.Params.SeenWithin)

		if e != nil {
			return server.GetRoster400JSONResponse(
				server.Error{
					Error: validation.New(
						"seen_within is not a duration (24h, 7d): %s",
						*r.Params.SeenWithin,
					).Error(),
				},
			), nil
		}

		window = time.Duration(d)
	}

	entries, e := s.roster.Roster(window)

	if e != nil {
		if validation.Is(e) {
			return server.GetRoster400JSONResponse(
				server.Error{Error: e.Error()},
			), nil
		}

		return server.GetRoster500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	result := server.GetRoster200JSONResponse{
		Entries: make([]server.RosterEntry, 0, len(entries)),
	}

	for _, v := range entries {
		result.Entries = append(
			result.Entries,
			server.RosterEntry{
				Callsign:   v.Callsign,
				Identifier: v.Identifier,
				LastSeen:   v.LastSeen,
				OnRoster:   v.OnRoster,
			},
		)
	}

	return result, nil
}
