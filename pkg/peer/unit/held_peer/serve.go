package held_peer

import "context"

func (p *Peer) Serve(x context.Context) error {
	p.served <- struct{}{}

	select {
	case <-x.Done():
		return nil
	case e := <-p.end:
		return e
	}
}
