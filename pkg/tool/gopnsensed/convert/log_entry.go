package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/log_entry"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func LogEntry(v *log_entry.Entry) *server.LogEntry {
	return &server.LogEntry{
		Timestamp:       v.Timestamp,
		Interface:       v.Interface,
		Action:          v.Action,
		Direction:       v.Direction,
		Protocol:        v.Protocol,
		Source:          v.Source,
		SourcePort:      v.SourcePort,
		Destination:     v.Destination,
		DestinationPort: v.DestinationPort,
		RuleNumber:      v.RuleNumber,
		RuleIdentifier:  v.RuleIdentifier,
		Label:           v.Label,
	}
}
