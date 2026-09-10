package goclaude

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func printDeleteReceipt(r *client.DeleteReceiptResponse) {
	console.Format("deleted %s (%s)\n", r.Identifier, r.Name)
	rows := [][]any{
		{"event", r.Events},
		{"event metadata", r.EventMetadata},
		{"completion", r.Completions},
		{"summary", r.Summaries},
		{"label", r.Labels},
		{"pulse", r.Pulses},
		{"context load", r.ContextLoads},
		{"tracker state", r.TrackerStates},
		{"queue", r.Queue},
		{"notification", r.Notifications},
	}

	for _, row := range rows {
		console.Format("  %-16s %d\n", row[0], row[1])
	}

	if r.Transcript != nil {
		console.Format("  %-16s %s\n", "transcript", *r.Transcript)
	}

	if r.Sources == nil {
		return
	}

	for _, path := range *r.Sources {
		console.Format("  %-16s %s\n", "source", path)
	}
}
