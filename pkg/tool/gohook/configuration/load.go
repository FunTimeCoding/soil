package configuration

import (
	"fmt"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/markup"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"path/filepath"
	"slices"
)

func Load(root string) *Configuration {
	path := system.FirstFile(
		filepath.Join(root, constant.RootFile),
		filepath.Join(root, constant.ToolFile),
	)

	if path == "" {
		panic(
			fmt.Sprintf(
				constant.MissingConfiguration,
				constant.RootFile,
				constant.ToolFile,
				root,
			),
		)
	}

	result := New(root, path)
	markup.MustDecode(system.ReadFileUnsafe(path), &result.Hooks)

	for hook, jobs := range result.Hooks {
		if !slices.Contains(gitConstant.Hooks, hook) {
			panic(
				fmt.Sprintf(
					constant.UnknownHook,
					hook,
					path,
					join.CommaSpace(gitConstant.Hooks),
				),
			)
		}

		for i, j := range jobs {
			if j == nil || j.Run == "" {
				panic(fmt.Sprintf(constant.MissingRun, i+1, hook))
			}
		}
	}

	return result
}
