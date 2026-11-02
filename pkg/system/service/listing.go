package service

import "github.com/funtimecoding/soil/pkg/system/run"

func listing(
	name string,
	a ...string,
) (string, error) {
	r := run.New().NoPanic()
	r.Start(append([]string{name}, a...)...)

	return r.OutputString, r.Error
}
