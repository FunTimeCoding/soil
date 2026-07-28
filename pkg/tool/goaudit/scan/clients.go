package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"sort"
)

func Clients(
	v *virtual_file_system.System,
	repository string,
	configuration *Configuration,
) []*Client {
	result := findClients(
		v,
		constant.PackageDirectory,
		constant.PackageDirectory,
		repository,
		configuration,
	)
	sort.Slice(
		result,
		func(
			i int,
			j int,
		) bool {
			if result[i].Repo != result[j].Repo {
				return result[i].Repo < result[j].Repo
			}

			return result[i].Path < result[j].Path
		},
	)

	return result
}
