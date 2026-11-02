package go_mod

import (
	"github.com/funtimecoding/soil/pkg/console"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func recoverDeadTag(
	name string,
	stderr string,
	skipProxy bool,
) *run.Run {
	mod, version := ParseDeadTag(stderr)

	if mod == "" {
		return nil
	}

	console.Format("Dead tag: %s@%s - recovering\n", mod, version)
	dropRequire(mod)
	cleanSum(mod, version)
	r := run.New()
	r.Panic = false

	if skipProxy {
		r.Environment(soil.Proxy, soil.Direct)
	}

	r.Start(soil.Go, soil.Get, name)

	return r
}
