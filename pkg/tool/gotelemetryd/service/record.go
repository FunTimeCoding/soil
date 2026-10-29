package service

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model/usage_event"
)

func (s *Service) Record(e *record.Record) {
	u := usage_event.New()
	u.Tool = e.Tool
	u.Surface = e.Surface
	u.Actor = e.Actor
	u.Outcome = e.Outcome
	u.Kind = e.Kind

	if len(e.Detail) > 0 {
		encoded, marshalError := json.Marshal(e.Detail)

		if marshalError == nil {
			u.Detail = new(string(encoded))
		}
	}

	if f := s.store.Create(u); f != nil {
		s.reporter.CaptureException(f)
	}
}
