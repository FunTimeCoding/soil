package resolve_result

import "github.com/funtimecoding/soil/pkg/tool/goclauded/types/match"

func New(matches []*match.Match) *Result {
	return &Result{Matches: matches}
}
