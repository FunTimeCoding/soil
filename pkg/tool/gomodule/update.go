package gomodule

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/go_mod"
	"github.com/funtimecoding/soil/pkg/runtime"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/terminal"
)

func update(
	t *terminal.Terminal,
	exclusives []string,
	downgrades []string,
	continueOnError bool,
) {
	if len(exclusives) > 0 {
		console.Format("Exclusive matches: %s\n", join.Comma(exclusives))
	}

	if len(downgrades) > 0 {
		console.Format("Downgrades: %s\n", join.Comma(downgrades))
	}

	go_mod.UpdateDirectDependencies(exclusives, continueOnError)
	go_mod.DowngradeDependencies(downgrades)
	go_mod.Tidy()
	v := runtime.ExecutableVersion()

	if v == nil {
		t.Exitf("could not determine Go version\n")

		return
	}

	system.Run(
		constant.Go,
		constant.Mod,
		constant.Edit,
		key_value.Equals(constant.VersionArgument, v.String()),
	)
	updateContainerFile(v)
	updateGitLabFile(v.String())
}
