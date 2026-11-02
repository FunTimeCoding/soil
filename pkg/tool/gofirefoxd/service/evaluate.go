package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	firefoxConstant "github.com/funtimecoding/soil/pkg/firefox/constant"
	"github.com/funtimecoding/soil/pkg/firefox/evaluation"
)

func (s *Service) Evaluate(
	identifier int,
	expression string,
) (*evaluation.Evaluation, error) {
	if expression == "" {
		return nil, validation.New(firefoxConstant.ExpressionRequired)
	}

	return s.client.Evaluate(identifier, expression)
}
