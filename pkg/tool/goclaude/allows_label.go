package goclaude

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/settings_file"
	"os"
)

func allowsLabel(path string) bool {
	raw, e := os.ReadFile(path)

	if e != nil {
		return false
	}

	var s settings_file.File

	if f := json.Unmarshal(raw, &s); f != nil {
		return false
	}

	for _, rule := range s.Permissions.Allow {
		if constant.LabelToolAllowance.MatchString(rule) {
			return true
		}
	}

	return false
}
