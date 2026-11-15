package switchable_peer

import "sync"

type Peer struct {
	mutex   sync.Mutex
	failure error
	probes  int
}
