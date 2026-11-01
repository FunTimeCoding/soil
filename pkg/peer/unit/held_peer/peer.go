package held_peer

import "sync"

type Peer struct {
	mutex   sync.Mutex
	refusal []error
	opened  int
	closed  int
	end     chan error
}
