package unit

import "sync"

type FixedEmbedder struct {
	mutex   sync.Mutex
	failure error
	calls   int
}
