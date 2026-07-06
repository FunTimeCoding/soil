package system

import "github.com/funtimecoding/soil/pkg/system/run"

func RunDirectory(
	directory string,
	s ...string,
) string {
	r := run.New()
	r.Directory = directory

	return r.Start(s...)
}
