package unit

import "sync"

type SwitchablePeer struct {
	mutex   sync.Mutex
	failure error
	probes  int
}
