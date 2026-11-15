package failing_dialer

import "sync"

type Dialer struct {
	failure  error
	mutex    sync.Mutex
	attempts int
}
