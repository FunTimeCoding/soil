package connector

import "time"

func NewTarget(
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
