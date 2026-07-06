package go_mod

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func DowngradeDependencies(v []string) {
	for _, e := range v {
		fmt.Printf("Downgrade: %s\n", e)
		system.Run(constant.Go, constant.Get, e)
	}

	if len(v) > 0 {
		Tidy()
	}
}
