package go_mod

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func DowngradeDependencies(v []string) {
	for _, e := range v {
		console.Format("Downgrade: %s\n", e)
		system.Run(constant.Go, constant.Get, e)
	}

	if len(v) > 0 {
		Tidy()
	}
}
