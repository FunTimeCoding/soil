package amtool

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/maps"
	alertmanager "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
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
		alertmanager.AmtoolPath,
	)
	tool := join.Absolute(base, alertmanager.AmtoolConfiguration)

	if !run.CommandExists(alertmanager.AmtoolCommand) {
		fmt.Println(
			"amtool missing: go install github.com/prometheus/alertmanager/cmd/amtool@latest",
		)

		return
	}

	if !system.FileExists(tool) {
		fmt.Printf("Missing: %s\n", tool)

		return
	}

	active := Read(base, alertmanager.AmtoolConfiguration)
	files := system.Files(base)
	var contexts []string
	locatorByContext := make(map[string]string)

	for _, f := range files {
		if false {
			fmt.Printf("File: %s\n", f)
		}

		name, _ := key_value.Dot(f)

		if strings.HasPrefix(name, alertmanager.AmtoolConfigurationPrefix) {
			context := strings.TrimPrefix(
				name,
				alertmanager.AmtoolConfigurationPrefix,
			)

			if !slices.Contains(contexts, context) {
				contexts = append(contexts, context)
				c := Read(base, f)
				locatorByContext[context] = c.Locator
			}
		}
	}

	if len(contexts) == 0 {
		fmt.Println("No contexts found")

		return
	}

	if selected == "" {
		for _, k := range maps.StringKeys(locatorByContext) {
			v := locatorByContext[k]

			if v == active.Locator {
				fmt.Printf("* %s\n", k)
			} else {
				fmt.Printf("  %s\n", k)
			}
		}

		return
	}

	if !slices.Contains(contexts, selected) {
		fmt.Printf("Context not found: %s\n", selected)

		return
	}

	if active.Locator == locatorByContext[selected] {
		fmt.Printf("Already active: %s\n", selected)

		return
	}

	system.CopyFile(
		join.Absolute(
			base,
			fmt.Sprintf(
				"%s%s.yml",
				alertmanager.AmtoolConfigurationPrefix,
				selected,
			),
		),
		tool,
	)
	fmt.Printf("Now active: %s\n", selected)
}
