package system

import "github.com/funtimecoding/soil/pkg/system/join"

func ParentDirectory(levels int) string {
	var dots []string

	for i := 0; i < levels; i++ {
		dots = append(dots, "..")
	}

	return join.Absolute(append([]string{WorkDirectory()}, dots...)...)
}
