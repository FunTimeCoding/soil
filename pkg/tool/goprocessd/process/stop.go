package process

import (
	"golang.org/x/sys/unix"
	"time"
)

func (p *Process) Stop() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.handle == nil {
		return nil
	}

	p.stoppedBySupervisor = true
	e := terminateGroup(p.handle.Pid())

	if e != nil && !Exited(e) {
		return e
	}

	timeout := time.AfterFunc(
		10*time.Second,
		func() {
			p.mutex.Lock()
			defer p.mutex.Unlock()

			if p.handle != nil {
				e := unix.Kill(-1*p.handle.Pid(), unix.SIGKILL)

				if e != nil {
					return
				}
			}
		},
	)
	p.condition.Wait()
	timeout.Stop()

	return nil
}
