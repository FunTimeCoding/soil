package pool

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/inventory/constant"
)

func (p *Pool[T, P, C]) Client(instance string) (C, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if c, okay := p.clients[instance]; okay {
		return c, nil
	}

	var zero C
	i, okay := p.inventory.Find(instance)

	if !okay {
		return zero, not_found.New(constant.InstanceKind, instance)
	}

	result, e := p.build(i)

	if e != nil {
		return zero, e
	}

	p.clients[instance] = result

	return result, nil
}
