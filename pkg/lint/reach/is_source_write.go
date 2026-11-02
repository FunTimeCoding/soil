package reach

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/reach/transcript"
	"strings"
)

func isSourceWrite(b *transcript.Block) bool {
	if !strings.HasPrefix(b.Name, constant.SourceToolPrefix) ||
		b.Name == constant.SourceUseModule {
		return false
	}

	var input map[string]any

	if json.Unmarshal(b.Input, &input) != nil {
		return false
	}

	dryRun, _ := input[constant.DryRunInput].(bool)

	return !dryRun
}
