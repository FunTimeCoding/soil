package failure

import "github.com/funtimecoding/soil/pkg/runtime"

func (f *Failure) Caller() error {
	name := runtime.CallerName(1)
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.errors[name]
}
