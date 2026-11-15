package mock_embedder

import "sync"

type Embedder struct {
	mutex     sync.Mutex
	failure   error
	embedding error
	refused   string
	inputs    []string
	calls     int
	closed    int
}
