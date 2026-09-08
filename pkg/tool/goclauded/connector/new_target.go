package connector

import "time"

func NewTarget(
	identifier string,
	name string,
	timestamp time.Time,
	labels map[string]string,
) *Target {
	if labels == nil {
		labels = map[string]string{}
	}

	return &Target{
		Identifier: identifier,
		Name:       name,
		Timestamp:  timestamp,
		Labels:     labels,
	}
}
