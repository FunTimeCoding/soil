package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
)

func (s *Summaries) holds(c *fact.Condition) bool {
	switch c.Kind {
	case "":
		return true
	case constant.DisposesCondition:
		return s != nil && s.disposes[c.Callee][c.Argument]
	case constant.ReleasesCondition:
		return s != nil && s.releases[c.Holder][c.Field]
	case constant.WrapsCondition:
		return s != nil && s.wraps[c.Callee][c.Argument]
	}

	return false
}
