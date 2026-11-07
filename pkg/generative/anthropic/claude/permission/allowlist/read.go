package allowlist

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/allowlist/setting"
	"os"
)

func Read(path string) ([]*Rule, error) {
	b, e := os.ReadFile(path)

	if e != nil {
		return nil, e
	}

	var f setting.File

	if g := json.Unmarshal(b, &f); g != nil {
		return nil, g
	}

	if f.Permissions == nil {
		return nil, nil
	}

	return Parse(f.Permissions.Allow), nil
}
