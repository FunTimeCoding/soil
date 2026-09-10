package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/receipt"
)

func deleteReceipt(r *receipt.Receipt) *server.DeleteReceiptResponse {
	result := &server.DeleteReceiptResponse{
		Identifier:    r.Identifier,
		Name:          r.Name,
		Events:        int(r.Events),
		EventMetadata: int(r.EventMetadata),
		Completions:   int(r.Completions),
		Summaries:     int(r.Summaries),
		Labels:        int(r.Labels),
		Pulses:        int(r.Pulses),
		ContextLoads:  int(r.ContextLoads),
		TrackerStates: int(r.TrackerStates),
		Queue:         int(r.Queue),
		Notifications: int(r.Notifications),
	}

	if r.Transcript != "" {
		result.Transcript = &r.Transcript
	}

	if len(r.Sources) > 0 {
		result.Sources = &r.Sources
	}

	return result
}
