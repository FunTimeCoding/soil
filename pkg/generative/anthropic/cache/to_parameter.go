package cache

import (
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/constant"
)

func ToParameter(m constant.Mode) anthropic.CacheControlEphemeralParam {
	p := anthropic.NewCacheControlEphemeralParam()

	if m == constant.OneHour {
		p.TTL = anthropic.CacheControlEphemeralTTLTTL1h
	}

	return p
}
