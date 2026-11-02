package reach

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"mvdan.cc/sh/v3/syntax"
)

func redirectTargets(s *syntax.Stmt) []*syntax.Word {
	var result []*syntax.Word

	for _, r := range s.Redirs {
		if r.Op != syntax.RdrOut && r.Op != syntax.AppOut ||
			r.N != nil ||
			r.Word == nil {
			continue
		}

		if v, okay := literalValue(r.Word); okay && v == constant.NullDevice {
			continue
		}

		result = append(result, r.Word)
	}

	return result
}
