package unit

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/port_forward"

func stopped(f *port_forward.Forward) bool {
	select {
	case <-f.Stopping():
		return true
	default:
		return false
	}
}
