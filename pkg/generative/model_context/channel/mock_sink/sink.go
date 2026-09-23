package mock_sink

import "sync"

type Sink struct {
	mutex sync.Mutex
	event []Event
}
