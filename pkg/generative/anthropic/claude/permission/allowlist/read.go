package allowlist

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/allowlist/setting"

func Read(path string) ([]*Rule, error) {
	f, e := setting.Read(path)

	if e != nil {
		return nil, e
	}

	if f.Permissions == nil {
		return nil, nil
	}

	return Parse(f.Permissions.Allow), nil
}
