package go_mod

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func recoverDeadTag(
	name string,
	stderr string,
	skipProxy bool,
) *run.Run {
	mod, version := parseDeadTag(stderr)

	if mod == "" {
		return nil
	}

	fmt.Printf("Dead tag: %s@%s - recovering\n", mod, version)
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
