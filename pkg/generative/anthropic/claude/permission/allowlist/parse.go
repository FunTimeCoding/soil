package allowlist

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"strings"
)

func Parse(entries []string) []*Rule {
	var result []*Rule

	for _, entry := range entries {
		if !strings.HasPrefix(entry, constant.RulePrefix) ||
			!strings.HasSuffix(entry, constant.RuleSuffix) {
			continue
		}

		body := strings.TrimSuffix(
			strings.TrimPrefix(entry, constant.RulePrefix),
			constant.RuleSuffix,
		)
		prefix := false

		for _, wildcard := range []string{
			constant.RuleColonWildcard,
			constant.RuleSpaceWildcard,
		} {
			if strings.HasSuffix(body, wildcard) {
				body = strings.TrimSuffix(body, wildcard)
				prefix = true

				break
			}
		}

		words := strings.Fields(body)

		if len(words) == 0 || strings.ContainsAny(body, "*") {
			continue
		}

		result = append(result, New(words, prefix))
	}

	return result
}
