package fetch

import (
	"github.com/funtimecoding/soil/pkg/monitor"
	"github.com/funtimecoding/soil/pkg/monitor/item"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func Run(c string) []*item.Item {
	if run.CommandExists(split.Space(c)[0]) {
		if m := monitor.Run(c); m != nil {
			return m.Items
		}
	}

	return nil
}
