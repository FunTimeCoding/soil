package main

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/jenkins"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func main() {
	j := jenkins.NewEnvironment()

	if false {
		plugins := j.Plugins()
		fmt.Printf("Count: %d\n", len(plugins))

		for _, p := range plugins {
			fmt.Printf("%s %s\n", p.ShortName, p.Version)

			if !p.Enabled {
				fmt.Println("    DISABLED")
			}
		}
	}

	if false {
		plugins := j.Basic().Get(
			"/pluginManager/api/json?tree=plugins[shortName,version,hasUpdate,enabled]",
		)
		fmt.Println(plugins)
	}

	if true {
		root := join.Absolute(
			system.WorkDirectory(),
			constant.Temporary,
			"jenkins",
		)
		system.EnsurePathExists(root)

		for _, o := range j.Jobs() {
			name := o.GetName()
			fmt.Println(name)

			if false {
				b := j.Basic().Get(
					fmt.Sprintf("job/%s/config.xml", name),
				)
				system.SaveFile(
					join.Absolute(root, fmt.Sprintf("%s.xml", name)),
					b,
				)
			}
		}
	}

	if false {
		j.JobsJSON()
	}
}
