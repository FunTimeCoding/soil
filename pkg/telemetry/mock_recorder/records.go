package mock_recorder

import (
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	"slices"
)

func (r *Recorder) Records() []*record.Record {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return slices.Clone(r.calls)
}
