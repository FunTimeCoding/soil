package main

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/jenkins"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func main() {
	j := jenkins.NewEnvironment()

	if false {
		plugins := j.Plugins()
		console.Format("Count: %d\n", len(plugins))

		for _, p := range plugins {
			console.Format("%s %s\n", p.ShortName, p.Version)

			if !p.Enabled {
				console.Line("    DISABLED")
			}
		}
	}

	if false {
		plugins := j.Basic().Get(
			"/pluginManager/api/json?tree=plugins[shortName,version,hasUpdate,enabled]",
		)
		console.Line(plugins)
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
			console.Line(name)

			if false {
				b := j.Basic().Get(fmt.Sprintf("job/%s/config.xml", name))
				system.SaveFile(
					join.Absolute(root, fmt.Sprintf("%s.xml", name)),
					b,
				)
			}
		}
	}

	if false {
		j.JobsNotation()
	}
}
