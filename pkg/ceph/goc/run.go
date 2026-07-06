package goc

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/system/join"
	"slices"
)

func Run(
	selected string,
	verbose bool,
) {
	file := environment.Required(ConfigurationEnvironment)

	if verbose {
		fmt.Printf("File: %s\n", file)
	}

	base := join.Absolute(system.Home(), constant.ConfigurationPath, CephPath)

	if verbose {
		fmt.Printf("Base: %s\n", base)
	}

	active := split.Slash(file)[5]

	if verbose {
		fmt.Printf("Active: %s\n", active)
	}

	directories := system.Directories(base)

	if selected == "" {
		for _, d := range directories {
			if d == active {
				fmt.Printf("* %s\n", d)
			} else {
				fmt.Printf("  %s\n", d)
			}
		}

		return
	}

	if !slices.Contains(directories, selected) {
		fmt.Printf("Unexpected: %s\n", selected)

		return
	}

	name := configurationName(base, selected)
	newConfiguration := join.Absolute(base, selected, clientConfiguration)
	newArgument := fmt.Sprintf(
		"-n %s --keyring=%s",
		fmt.Sprintf("client.%s", name),
		join.Absolute(
			base,
			selected,
			fmt.Sprintf("ceph.client.%s.keyring", name),
		),
	)

	if verbose {
		fmt.Printf("newConfiguration: %s\n", newConfiguration)
		fmt.Printf("newArgument: %s\n", newArgument)
	}

	environment.SetTerminal(ConfigurationEnvironment, newConfiguration)
	environment.SetTerminal(ArgumentEnvironment, newArgument)
}
