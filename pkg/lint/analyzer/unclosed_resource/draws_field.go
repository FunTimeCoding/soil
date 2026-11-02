package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
)

func (s *Summaries) drawsField(c *fact.Condition) bool {
	if c.Field == constant.AnyField {
		return len(s.draws[c.Callee][c.Argument]) > 0
	}

	return s.draws[c.Callee][c.Argument][c.Field]
}
