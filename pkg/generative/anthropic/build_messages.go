package anthropic

import (
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/cache"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/message"
)

func buildMessages(
	v []*message.Message,
	m constant.Mode,
) []anthropic.MessageParam {
	result := message.ToParameterSlice(v)

	if m != constant.None && len(result) >= 2 {
		result[len(result)-2] = v[len(v)-2].ToParameterCached(
			cache.ToParameter(m),
		)
	}

	return result
}
