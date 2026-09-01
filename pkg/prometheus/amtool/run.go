package amtool

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/maps"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/system/run"
	"slices"
	"strings"
)

func Run(selected string) {
	base := join.Absolute(
		system.Home(),
		constant.ConfigurationPath,
		prometheus.AmtoolPath,
	)
	tool := join.Absolute(base, prometheus.AmtoolConfiguration)

	if !run.CommandExists(prometheus.AmtoolCommand) {
		console.Line(
			"amtool missing: go install github.com/prometheus/alertmanager/cmd/amtool@latest",
		)

		return
	}

	if !system.FileExists(tool) {
		console.Format("Missing: %s\n", tool)

		return
	}

	active := Read(base, prometheus.AmtoolConfiguration)
	files := system.Files(base)
	var contexts []string
	locatorByContext := make(map[string]string)

	for _, f := range files {
		if false {
			console.Format("File: %s\n", f)
		}

		name, _ := key_value.Dot(f)

		if strings.HasPrefix(name, prometheus.AmtoolConfigurationPrefix) {
			context := strings.TrimPrefix(
				name,
				prometheus.AmtoolConfigurationPrefix,
			)

			if !slices.Contains(contexts, context) {
				contexts = append(contexts, context)
				c := Read(base, f)
				locatorByContext[context] = c.Locator
			}
		}
	}

	if len(contexts) == 0 {
		console.Line("No contexts found")

		return
	}

	if selected == "" {
		for _, k := range maps.StringKeys(locatorByContext) {
			v := locatorByContext[k]

			if v == active.Locator {
				console.Format("* %s\n", k)
			} else {
				console.Format("  %s\n", k)
			}
		}

		return
	}

	if !slices.Contains(contexts, selected) {
		console.Format("Context not found: %s\n", selected)

		return
	}

	if active.Locator == locatorByContext[selected] {
		console.Format("Already active: %s\n", selected)

		return
	}

	system.CopyFile(
		join.Absolute(
			base,
			fmt.Sprintf(
				"%s%s.yml",
				prometheus.AmtoolConfigurationPrefix,
				selected,
			),
		),
		tool,
	)
	console.Format("Now active: %s\n", selected)
}
