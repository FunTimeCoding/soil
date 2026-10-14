package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/types"
)

func (s *Summaries) takes(
	f *types.Func,
	argument int,
) bool {
	key := functionKey(f)

	return s.holds(argumentCondition(f, argument)) ||
		s.holds(fact.NewWrapCondition(key, argument))
}
