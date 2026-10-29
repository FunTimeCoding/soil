package mock_system

import (
	"github.com/funtimecoding/soil/pkg/system/run/stream/mock_stream"
	"sync"
)

type System struct {
	mutex    sync.Mutex
	outputs  map[string]string
	failures map[string]error
	stream   *mock_stream.Stream
}
