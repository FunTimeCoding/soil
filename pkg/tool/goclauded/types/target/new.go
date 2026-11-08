package target

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/label_value"
	"time"
)

func New(
	identifier string,
	name string,
	lastSeen time.Time,
	labels map[string]string,
) *Target {
	sets := map[string][]string{}

	for key, value := range labels {
		if values := label_value.Parse(value); len(values) > 0 {
			sets[key] = values
		}
	}

	return &Target{
		Identifier: identifier,
		Name:       name,
		LastSeen:   lastSeen,
		Labels:     sets,
	}
}
