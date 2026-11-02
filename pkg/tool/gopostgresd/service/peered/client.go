package peered

import "github.com/funtimecoding/soil/pkg/postgres/face"

func (p *Peered) Client() face.Client {
	return p.client
}
