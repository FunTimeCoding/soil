package mock_recorder

import (
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	"sync"
)

type Recorder struct {
	mutex sync.Mutex
	calls []*record.Record
}
