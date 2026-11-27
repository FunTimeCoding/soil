package service

import (
	browserConstant "github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/firefox/evaluation"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func (s *Service) Evaluate(
	identifier int,
	expression string,
) (*evaluation.Evaluation, error) {
	if expression == "" {
		return nil, validation.New(browserConstant.FirefoxExpressionRequired)
	}

	return s.client.Evaluate(identifier, expression)
}
