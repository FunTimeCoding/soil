package log_entry

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func New(v response.LogEntry) *Entry {
	return &Entry{
		Timestamp:       v.Timestamp,
		Interface:       v.Interface,
		Action:          v.Action,
		Direction:       v.Direction,
		Protocol:        v.ProtocolName,
		Source:          v.Source,
		SourcePort:      v.SourcePort,
		Destination:     v.Destination,
		DestinationPort: v.DestinationPort,
		RuleNumber:      v.RuleNumber,
		RuleIdentifier:  v.RuleIdentifier,
		Label:           v.Label,
	}
}
