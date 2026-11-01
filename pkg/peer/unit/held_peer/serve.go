package held_peer

import "context"

func (p *Peer) Serve(x context.Context) error {
	select {
	case <-x.Done():
		return nil
	case e := <-p.ending():
		return e
	}
}
