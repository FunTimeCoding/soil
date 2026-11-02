package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"strconv"
)

func (s *Summaries) relays(c *fact.Condition) bool {
	sink, e := strconv.Atoi(c.Field)

	if e != nil {
		return false
	}

	for field := range s.yields[c.Callee][c.Argument] {
		if s.draws[c.Holder][sink][field] {
			return true
		}
	}

	return false
}
