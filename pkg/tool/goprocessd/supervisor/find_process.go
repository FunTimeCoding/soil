package supervisor

import "github.com/funtimecoding/soil/pkg/tool/goprocessd/process"

func (s *Supervisor) findProcess(name string) *process.Process {
	for _, p := range s.snapshotProcesses() {
		if p.Name == name {
			return p
		}
	}

	return nil
}
