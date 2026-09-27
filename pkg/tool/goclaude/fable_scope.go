package goclaude

import "github.com/funtimecoding/soil/pkg/generative/constant"

func fableScope(limits *statusLineRateLimits) *statusLineModelScoped {
	if limits == nil {
		return nil
	}

	for _, m := range limits.ModelScoped {
		if m.DisplayName == constant.UsageScopeFable {
			return &m
		}
	}

	return nil
}
