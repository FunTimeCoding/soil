package configuration

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/markup"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	tool "github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"path/filepath"
	"slices"
)

func Load(root string) *Configuration {
	path := system.FirstFile(
		filepath.Join(root, tool.RootFile),
		filepath.Join(root, tool.ToolFile),
	)

	if path == "" {
		panic(
			fmt.Sprintf(
				tool.MissingConfiguration,
				tool.RootFile,
				tool.ToolFile,
				root,
			),
		)
	}

	result := New(root, path)
	markup.MustDecode(system.ReadFileUnsafe(path), &result.Hooks)

	for hook, jobs := range result.Hooks {
		if !slices.Contains(constant.Hooks, hook) {
			panic(
				fmt.Sprintf(
					tool.UnknownHook,
					hook,
					path,
					join.CommaSpace(constant.Hooks),
				),
			)
		}

		for i, j := range jobs {
			if j == nil || j.Run == "" {
				panic(fmt.Sprintf(tool.MissingRun, i+1, hook))
			}
		}
	}

	return result
}
