package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_result"
)

func AffectedSystems(v *affected_result.Result) *server.AffectedSystems {
	systems := make([]server.AffectedSystem, len(v.Rows))

	for i, r := range v.Rows {
		systems[i] = server.AffectedSystem{
			System:    r.System,
			Coverage:  r.Coverage,
			Priority:  r.Priority,
			Source:    r.Source,
			Installed: r.Installed,
			Fixed:     r.Fixed,
			Candidate: r.Candidate,
			Kind:      r.Kind,
		}

		if len(r.Labels) > 0 {
			systems[i].Labels = &r.Labels
		}
	}

	result := &server.AffectedSystems{
		Identifier: v.Identifier,
		Count:      AffectedCount(v.Count),
		Systems:    systems,
		Severity:   optional(v.Severity),
	}

	if v.Severity != "" {
		result.Score = &v.Score
	}

	return result
}
