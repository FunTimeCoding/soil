package runner

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/output"
	"strings"
)

func (r *Runner) needsUpgrade(s string) bool {
	for _, l := range split.NewLine(s) {
		var e output.DiagnosticEntry

		if json.Unmarshal([]byte(l), &e) != nil {
			continue
		}

		if strings.Contains(
			e.Diagnostic.Detail,
			"terraform init -upgrade",
		) {
			return true
		}
	}

	return false
}
