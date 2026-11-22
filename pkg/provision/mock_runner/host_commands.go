package mock_runner

import "github.com/funtimecoding/soil/pkg/provision/generated/server"

func (r *Runner) HostCommands() []server.HostCommandRequest {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return append([]server.HostCommandRequest{}, r.hostCommands...)
}
