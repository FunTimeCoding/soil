package cache

import (
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func ToParameter(m constant.AnthropicMode) anthropic.CacheControlEphemeralParam {
	p := anthropic.NewCacheControlEphemeralParam()

	if m == constant.AnthropicModeOneHour {
		p.TTL = anthropic.CacheControlEphemeralTTLTTL1h
	}

	return p
}
