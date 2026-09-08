package service

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func (s *Service) UtilizationFallback() bool {
	return utilization.Supported() &&
		environment.Exists(constant.AnthropicBaseLinkEnvironment)
}
