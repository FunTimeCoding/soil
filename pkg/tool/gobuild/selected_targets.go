package gobuild

import systemConstant "github.com/funtimecoding/soil/pkg/system/constant"

func selectedTargets(flags map[string]*bool) []string {
	var result []string

	for _, t := range systemConstant.SystemArchitectures {
		if *flags[t] {
			result = append(result, t)
		}
	}

	if len(result) == 0 {
		return systemConstant.SystemArchitectures
	}

	return result
}
