package held_peer

func New() *Peer {
	return &Peer{end: make(chan error, 1), served: make(chan struct{}, 16)}
}
