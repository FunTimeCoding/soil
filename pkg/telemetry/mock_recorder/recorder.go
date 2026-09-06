package mock_recorder

import "github.com/funtimecoding/soil/pkg/telemetry/record"

type Recorder struct {
	Calls []*record.Record
}
