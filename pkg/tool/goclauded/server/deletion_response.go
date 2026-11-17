package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/deletion_result"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func deletionResponse(r *deletion_result.Result) *server.DeletionResult {
	result := &server.DeletionResult{
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
		SearchEntries: int(r.SearchEntries),
	}

	if r.Transcript != "" {
		result.Transcript = &r.Transcript
	}

	if len(r.Sources) > 0 {
		result.Sources = &r.Sources
	}

	return result
}
