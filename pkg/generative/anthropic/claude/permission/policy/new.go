package policy

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/allowlist"

func New(
	rules []*allowlist.Rule,
	roots []string,
	base string,
) *Walk {
	result := &Walk{rules: rules, variables: map[string]string{}}

	for _, r := range roots {
		if resolved, okay := settle(r); okay {
			result.roots = append(result.roots, resolved)
		}
	}

	if resolved, okay := settle(base); okay {
		result.base = resolved
		result.origin = resolved
	}

	return result
}
