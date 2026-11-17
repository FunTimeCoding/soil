package transcript

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/allowlist"

func projectRules(
	cache map[string][]*allowlist.Rule,
	project string,
) ([]*allowlist.Rule, error) {
	if rules, okay := cache[project]; okay {
		return rules, nil
	}

	rules, e := allowlist.ReadAll(settingsPaths(project)...)

	if e != nil {
		return nil, e
	}

	cache[project] = rules

	return rules, nil
}
