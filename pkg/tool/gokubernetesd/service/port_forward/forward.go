package port_forward

import "sync"

type Forward struct {
	Identifier string
	Resource   string
	Namespace  string
	LocalPort  int
	TargetPort int
	stop       chan struct{}
	once       sync.Once
}
