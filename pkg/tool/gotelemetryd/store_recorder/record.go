package store_recorder

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

func (r *Recorder) Record(e *record.Record) {
	u := store.NewUsageEvent()
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

	if f := r.store.Create(u); f != nil {
		r.reporter.CaptureException(f)
	}
}
