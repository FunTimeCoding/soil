package go_mod

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/constant"
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
		r.Environment(constant.Proxy, constant.Direct)
	}

	r.Start(constant.Go, constant.Get, name)

	return r
}
