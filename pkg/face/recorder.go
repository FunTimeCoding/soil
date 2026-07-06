package face

import "github.com/funtimecoding/soil/pkg/telemetry/record"

type Recorder interface {
	Record(r *record.Record)
}
