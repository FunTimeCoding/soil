package loader

import "github.com/funtimecoding/soil/pkg/system"

func (l *Loader) Load(path string) {
	for _, n := range system.Files(path) {
		l.contents[n] = system.ReadFile(path, n)
	}
}
