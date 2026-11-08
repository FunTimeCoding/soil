package goclaude

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func CompactionPhase(event string) (client.CompactionPhase, bool) {
	switch event {
	case constant.PreCompactEvent:
		return client.Started, true
	case constant.PostCompactEvent:
		return client.Finished, true
	}

	return "", false
}
