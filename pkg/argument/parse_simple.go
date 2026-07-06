package argument

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func (i *Instance) ParseSimple() {
	errors.PanicOnError(i.flags.Parse(os.Args[1:]))
}
