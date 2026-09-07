package environment

import "sync"

type Environment struct {
	base     []string
	overlay  map[string]string
	exported map[string]struct{}
	deleted  map[string]struct{}
	mutex    sync.RWMutex
}
