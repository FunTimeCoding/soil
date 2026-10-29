package failure

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/runtime"
	"strings"
)

func (f *Failure) Set(
	method any,
	e error,
) {
	name := runtime.MethodName(method)

	if !strings.HasPrefix(name, f.owner) {
		errors.PanicOnError(
			unexpected.Format("%s is not a method of %s", name, f.owner),
		)
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.errors[name] = e
}
