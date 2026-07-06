package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"sort"
)

func Clients(
	v *virtual_file_system.System,
	repository string,
	configuration *Configuration,
) []*Client {
	result := findClients(
		v,
		"pkg",
		"pkg",
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
