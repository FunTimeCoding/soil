package anthropic

import (
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/cache"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func buildSystem(
	system string,
	m constant.AnthropicMode,
) []anthropic.TextBlockParam {
	if m == constant.AnthropicModeNone {
		return []anthropic.TextBlockParam{{Text: system}}
	}

	return []anthropic.TextBlockParam{
		{Text: system, CacheControl: cache.ToParameter(m)},
	}
}
