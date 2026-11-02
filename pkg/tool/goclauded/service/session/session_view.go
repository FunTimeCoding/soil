package session

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/session_view"
)

func (s *Service) SessionView(
	x context.Context,
	query string,
) (*session_view.View, error) {
	d, e := s.SessionDetail(query)

	if e != nil {
		return nil, e
	}

	result := session_view.New(d)

	if result.Labels, e = s.label.LabelsBySession(d.Identifier); e != nil {
		return nil, e
	}

	loads, f := s.transcript.ContextLoadsBySession(x, d.Identifier)

	if unavailable.Is(f) {
		result.LoadsHidden = f.Error()
	} else if f != nil {
		return nil, f
	}

	result.Loads = loads

	if result.Pulses, e = s.message.PulsesBySession(d.Identifier); e != nil {
		return nil, e
	}

	if result.Cruise, e = s.message.CruiseSetting(d.Identifier); e != nil {
		return nil, e
	}

	result.DeleteHash = s.DeleteHash(d.Identifier)

	return result, nil
}
