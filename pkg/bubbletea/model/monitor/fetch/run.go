package fetch

import (
	"github.com/funtimecoding/soil/pkg/monitor"
	"github.com/funtimecoding/soil/pkg/monitor/item"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func Run(c string) []*item.Item {
	if run.CommandExists(c) {
		if m := monitor.Run(c); m != nil {
			return m.Items
		}
	}

	return nil
}
