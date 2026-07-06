package system

import "github.com/funtimecoding/soil/pkg/system/run"

func RunError(s ...string) (string, error) {
	r := run.New().NoPanic()
	r.Start(s...)

	return r.OutputString, r.Error
}
