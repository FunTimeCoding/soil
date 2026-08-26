package packager

import "github.com/funtimecoding/soil/pkg/system/join"

func (p *Packager) scriptPath(name string) string {
	return join.Absolute(p.ConfigurationRoot, name)
}
