package runner

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/types/host_event"
)

func limit(parameters map[string]any) ([]string, bool) {
	events, okay := parameters[constant.DownstreamHosts].([]*host_event.Event)

	if !okay {
		return nil, false
	}

	var result []string

	for _, v := range events {
		if v.Action != constant.HostDelete {
			result = append(result, v.Host)
		}
	}

	return result, true
}
