package dictionary

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
)

func Read(file string) []*Category {
	f := system.Open(file)
	defer errors.LogClose(f)

	return Parse(f)
}
