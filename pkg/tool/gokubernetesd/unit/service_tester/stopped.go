package service_tester

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward"

func Stopped(f *port_forward.Forward) bool {
	select {
	case <-f.Stopping():
		return true
	default:
		return false
	}
}
