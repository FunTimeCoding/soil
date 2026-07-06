package build

import "github.com/funtimecoding/soil/pkg/system/constant"

func SystemArchitectures() []string {
	var result []string

	for _, operatingSystem := range constant.OperatingSystems {
		for _, architecture := range constant.Architectures {
			result = append(
				result,
				SystemArchitecture(operatingSystem, architecture),
			)
		}
	}

	return result
}
