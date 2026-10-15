package gobuild

import "github.com/funtimecoding/soil/pkg/system/constant"

func selectedTargets(flags map[string]*bool) []string {
	var result []string

	for _, t := range constant.SystemArchitectures {
		if *flags[t] {
			result = append(result, t)
		}
	}

	if len(result) == 0 {
		return constant.SystemArchitectures
	}

	return result
}
