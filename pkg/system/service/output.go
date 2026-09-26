package service

import "github.com/funtimecoding/soil/pkg/system/run"

func output(
	name string,
	a ...string,
) string {
	r := run.New()
	r.Panic = false
	r.Start(append([]string{name}, a...)...)

	return r.OutputString
}
