package tunnel

import "github.com/funtimecoding/soil/pkg/system/run/process"

type Tunnel struct {
	process   *process.Process
	started   chan struct{}
	listening chan struct{}
	stopped   chan struct{}
	Verbose   bool
	NoOutput  bool
	CleanStop bool
}
