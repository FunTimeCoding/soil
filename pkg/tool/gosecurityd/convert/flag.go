package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
)

func Flag(v *flag.Flag) *server.Flag {
	result := &server.Flag{
		Identifier: v.Subject,
		Note:       optional(v.Note),
		Severity:   optional(v.Severity),
		Mitigated:  v.Mitigated,
		Mitigation: optional(v.Mitigation),
		Flagged:    v.CreatedAt,
		Updated:    v.UpdatedAt,
	}

	if v.Severity != "" {
		result.Score = &v.Score
	}

	return result
}
