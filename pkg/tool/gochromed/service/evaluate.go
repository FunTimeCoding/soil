package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
)

func (s *Service) Evaluate(
	v *selector.Selector,
	expression string,
) (any, error) {
	if expression == "" {
		return nil, validation.New("expression is required")
	}

	t, e := s.resolveTab(v)

	if e != nil {
		return nil, e
	}

	var result any

	if f := s.source.Page(t.Identifier).EvaluatePromise(
		expression,
		&result,
	); f != nil {
		return nil, f
	}

	return result, nil
}
