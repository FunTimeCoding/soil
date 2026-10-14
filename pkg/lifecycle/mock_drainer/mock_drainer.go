package mock_drainer

import "sync"

type MockDrainer struct {
	Tag      string
	Log      *[]string
	hold     bool
	ignore   bool
	finished chan struct{}
	once     sync.Once
}
