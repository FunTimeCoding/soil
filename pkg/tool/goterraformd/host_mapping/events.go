package host_mapping

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/types/host_event"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/types/plan"
	"maps"
	"slices"
)

func (x *Index) Events(changes []*plan.ResourceChange) []*host_event.Event {
	merged := map[string]string{}

	for _, c := range changes {
		a := action(c.Change.Actions)

		if a == "" {
			continue
		}

		if _, named := x.attribute[c.Type]; named {
			if h := x.ownHost(c, a); h != "" {
				merged[h] = stronger(merged[h], a)
			}

			continue
		}

		for _, h := range x.hosts(configurationAddress(c.Address), map[string]bool{}) {
			merged[h] = stronger(merged[h], constant.HostUpdate)
		}
	}

	result := []*host_event.Event{}

	for _, h := range slices.Sorted(maps.Keys(merged)) {
		result = append(result, host_event.New(h, merged[h]))
	}

	return result
}
