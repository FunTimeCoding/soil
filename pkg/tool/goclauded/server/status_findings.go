package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/finding"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func statusFindings(findings []*finding.Finding) []server.FindingEntry {
	result := make([]server.FindingEntry, 0, len(findings))

	for _, i := range findings {
		entry := server.FindingEntry{Kind: i.Kind, Detail: i.Detail}

		if i.Subject != "" {
			entry.Subject = &i.Subject
		}

		result = append(result, entry)
	}

	return result
}
