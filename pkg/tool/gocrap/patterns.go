package gocrap

import crapConstant "github.com/funtimecoding/soil/pkg/crap/constant"

func patterns(arguments []string) []string {
	if len(arguments) == 0 {
		return []string{crapConstant.AllPackages}
	}

	return arguments
}
