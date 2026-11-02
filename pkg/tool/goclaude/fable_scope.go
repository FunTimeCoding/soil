package goclaude

import generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"

func fableScope(limits *StatusLineRateLimits) *StatusLineModelScoped {
	if limits == nil {
		return nil
	}

	for _, m := range limits.ModelScoped {
		if m.DisplayName == generativeConstant.UsageScopeFable {
			return &m
		}
	}

	return nil
}
