package target

import "time"

func New(
	identifier string,
	name string,
	lastSeen time.Time,
	labels map[string]string,
) *Target {
	if labels == nil {
		labels = map[string]string{}
	}

	return &Target{
		Identifier: identifier,
		Name:       name,
		LastSeen:   lastSeen,
		Labels:     labels,
	}
}
