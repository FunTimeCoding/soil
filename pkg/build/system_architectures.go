package build

import systemConstant "github.com/funtimecoding/soil/pkg/system/constant"

func SystemArchitectures() []string {
	var result []string

	for _, operatingSystem := range systemConstant.OperatingSystems {
		for _, architecture := range systemConstant.Architectures {
			result = append(
				result,
				SystemArchitecture(operatingSystem, architecture),
			)
		}
	}

	return result
}
