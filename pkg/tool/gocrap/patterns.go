package gocrap

import "github.com/funtimecoding/soil/pkg/crap/constant"

func patterns(arguments []string) []string {
	if len(arguments) == 0 {
		return []string{constant.AllPackages}
	}

	return arguments
}
