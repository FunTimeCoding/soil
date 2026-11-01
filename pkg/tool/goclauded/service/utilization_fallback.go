package service

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) UtilizationFallback() bool {
	return utilization.Supported() &&
		(environment.Exists(generativeConstant.AnthropicBaseLinkEnvironment) ||
			environment.Exists(constant.UsageEndpointEnvironment))
}
