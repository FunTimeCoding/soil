package process

import "time"

func (p *Process) StartedAt() time.Time {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.started
}
