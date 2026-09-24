package repository_tester

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func Command(
	directory string,
	arguments ...string,
) string {
	r := run.New()
	r.Directory = directory

	return r.Start(append([]string{constant.Command}, arguments...)...)
}
