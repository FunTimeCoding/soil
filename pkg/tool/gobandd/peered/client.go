package peered

import "github.com/funtimecoding/soil/pkg/band/face"

func (p *Peered) Client() face.BandClient {
	return p.client
}
