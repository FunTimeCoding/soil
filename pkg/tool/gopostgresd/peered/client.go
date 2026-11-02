package peered

import "github.com/funtimecoding/soil/pkg/postgres/face"

func (p *Peered) Client() face.PostgresClient {
	return p.client
}
