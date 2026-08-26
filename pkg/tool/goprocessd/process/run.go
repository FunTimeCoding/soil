package process

import (
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/system/writer"
)

func (p *Process) run(environment []string) bool {
	p.logger.StartGeneration()
	r := run.New()
	r.Panic = false
	r.SetEnvironment(environment)
	r.Writers(p.logger, p.logger)
	r.ProcessGroup()
	handle, e := r.TryOpen("/bin/sh", "-c", p.Command)

	if e != nil {
		writer.Print(p.logger, "Failed to start %s: %s\n", p.Name, e)

		return false
	}

	p.handle = handle
	p.stoppedBySupervisor = false
	p.mutex.Unlock()
	e = handle.Wait()
	p.mutex.Lock()
	p.condition.Broadcast()
	p.waitError = e
	p.handle = nil
	supervisorStopped := p.stoppedBySupervisor
	writer.Print(p.logger, "Terminating %s\n", p.Name)

	return supervisorStopped
}
