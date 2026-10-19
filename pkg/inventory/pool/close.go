package pool

import "github.com/funtimecoding/soil/pkg/inventory/face"

func (p *Pool[T, P, C]) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for name := range p.built {
		if v, okay := any(p.clients[name]).(face.Closer); okay {
			v.Close()
		}
	}

	clear(p.clients)
	clear(p.built)
}
