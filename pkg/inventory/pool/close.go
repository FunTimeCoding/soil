package pool

import "github.com/funtimecoding/soil/pkg/inventory/face"

func (p *Pool[T, P, C]) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, c := range p.clients {
		if v, okay := any(c).(face.Closer); okay {
			v.Close()
		}
	}

	clear(p.clients)
}
