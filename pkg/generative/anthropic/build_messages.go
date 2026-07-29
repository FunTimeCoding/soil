package anthropic

import (
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/cache"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/message"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func buildMessages(
	v []*message.Message,
	m constant.AnthropicMode,
) []anthropic.MessageParam {
	result := message.ToParameterSlice(v)

	if m != constant.AnthropicModeNone && len(result) >= 2 {
		result[len(result)-2] = v[len(v)-2].ToParameterCached(
			cache.ToParameter(m),
		)
	}

	return result
}
