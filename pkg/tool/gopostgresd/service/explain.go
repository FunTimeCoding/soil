package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
)

func (s *Service) Explain(
	x context.Context,
	instance string,
	sql string,
	analyze bool,
) (*result.Result, error) {
	prefix := constant.ExplainPrefix

	if analyze {
		prefix = constant.ExplainAnalyzePrefix
	}

	return s.guarded(x, instance, theirs, join.Space(prefix, sql))
}
