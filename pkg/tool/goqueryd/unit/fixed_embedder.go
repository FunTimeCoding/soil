package unit

import "sync"

type FixedEmbedder struct {
	mutex     sync.Mutex
	failure   error
	embedding error
	calls     int
	closed    int
}
