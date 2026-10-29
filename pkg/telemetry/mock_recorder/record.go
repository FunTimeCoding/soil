package mock_recorder

import "github.com/funtimecoding/soil/pkg/telemetry/record"

func (r *Recorder) Record(e *record.Record) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.calls = append(r.calls, e)
}
