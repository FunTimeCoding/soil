package service_tester

import "github.com/funtimecoding/soil/pkg/peer"

func (t *Tester) MemoryGuard() *peer.Guard {
	return t.memoryGuard
}
