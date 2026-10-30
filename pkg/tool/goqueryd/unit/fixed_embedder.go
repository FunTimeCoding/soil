package unit

import "sync"

type FixedEmbedder struct {
	mutex     sync.Mutex
	failure   error
	embedding error
	refused   string
	calls     int
	closed    int
}
