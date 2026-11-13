package layer

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func (l *Layer) touch(path string) {
	if _, seen := l.original[path]; seen {
		return
	}

	before, e := l.Read(path)

	if e != nil && !os.IsNotExist(e) {
		errors.PanicOnError(e)
	}

	l.original[path] = before
	l.order = append(l.order, path)
}
