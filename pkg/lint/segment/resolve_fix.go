package segment

import (
	"go/token"
	"go/types"
)

func ResolveFix(
	name string,
	segment string,
	applicable []string,
	o types.Object,
) string {
	scope := o.Parent()
	fix := ChooseFix(name, applicable)
	replacement := ReplaceSegment(name, segment, fix)

	if token.IsKeyword(replacement) {
		return ""
	}

	if scope == nil || !ScopeContains(scope, replacement) {
		return fix
	}

	if len(fix) > 1 {
		return ""
	}

	for _, a := range applicable {
		if len(a) == 1 && a != fix {
			candidate := ReplaceSegment(name, segment, a)

			if !ScopeContains(scope, candidate) {
				return a
			}
		}
	}

	multiCharacter := FirstMultiCharacter(applicable)

	if multiCharacter != "" {
		for _, r := range multiCharacter {
			letter := string(r)
			candidate := ReplaceSegment(name, segment, letter)

			if !ScopeContains(scope, candidate) {
				return letter
			}
		}
	}

	for c := byte('a'); c <= byte('z'); c++ {
		letter := string(c)
		candidate := ReplaceSegment(name, segment, letter)

		if !ScopeContains(scope, candidate) {
			return letter
		}
	}

	return ""
}
